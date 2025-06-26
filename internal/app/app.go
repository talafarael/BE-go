package app

import (
	"context"
	"errors"
	"fmt"
	"gin/docs"
	"gin/internal/app/controllers"
	"gin/internal/app/services"
	"gin/internal/infrastructure/config"
	"gin/internal/infrastructure/database/migrattion"
	"gin/internal/infrastructure/repository"
	"gin/internal/infrastructure/repository/postgres"
	check_auth_header "gin/pkg/checkAuthHeader"
	"gin/pkg/database"
	"gin/pkg/handler"
	"gin/pkg/hash"
	"gin/pkg/jwt"
	"gin/pkg/middleware"
	"gin/pkg/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db = make(map[string]string)

type App struct {
	config  *config.Config
	handler *handler.Handler
	srv     *server.Server
	store   database.Database
}

func NewApp(config config.Config) *App {
	store := database.NewGormDatabase(config.Connection).GetDB()
	repo := postgres.NewRepository(store)
	services := ConfigService(repo, &config)
	app := &App{
		config:  &config,
		handler: handler.NewHandler(*services),
		srv:     new(server.Server),
		store:   database.NewGormDatabase(config.Connection),
	}
	app.configureRouter()
	migrattion.RunMigrations(app.store)
	return app
}

func (app *App) Start() error {
	go func() {
		if err := app.srv.Run(app.config, app.handler); !errors.Is(err, http.ErrServerClosed) {
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := app.srv.Stop(ctx); err != nil {
		_ = fmt.Errorf("failed to stop server: %v", err)
	}
	return nil
}

func ConfigService(repo repository.Store, config *config.Config) *services.Service {
	hashService := hash.NewHashService()
	jwtService := jwt.NewJwtService(config.Secret)

	service := services.NewService(
		&services.ServiceOptions{
			Repo:        &repo,
			JwtService:  &jwtService,
			HashService: &hashService,
		})
	return &service
}

func ConfigMiddleware(repo repository.Store, config *config.Config) *middleware.AuthMiddleware {
	jwtService := jwt.NewJwtService(config.Secret)
	checkAuthHeader := check_auth_header.CheckAuthHeader{}

	authMiddleware := middleware.NewAuthMiddleware(
		&middleware.AuthMiddlewareOptions{
			CheckAuthHeader: &checkAuthHeader,
			JwtService:      &jwtService,
			Repo:            &repo,
		})
	return &authMiddleware
}

func (app *App) configureRouter() {
	router := app.handler.Routing()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	// swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	repo := postgres.NewRepository(app.store.GetDB())
	// Regiuster service
	services := ConfigService(repo, app.config)
	middleware := ConfigMiddleware(repo, app.config)

	controller := controllers.NewBaseController(services, middleware)
	controller.RegisterRoutes(router)
}

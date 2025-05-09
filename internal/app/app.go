package app

import (
	"context"
	"errors"
	"fmt"
	"gin/docs"
	"gin/internal/config"
	"gin/internal/controllers"
	"gin/internal/database/migrattion"
	"gin/internal/repository/postgres"
	"gin/internal/services"
	"gin/pkg/database"
	"gin/pkg/handler"
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
	app := &App{
		config:  &config,
		handler: handler.NewHandler(services.NewService(repo)),
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

func (app *App) configureRouter() {
	router := app.handler.Routing()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	repo := postgres.NewRepository(app.store.GetDB())

	// swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Regiuster service
	controller := controllers.NewBaseController(services.NewService(repo))
	controller.AddSingleController(controllers.NewUserController)
	controller.RegisterRoutes(router)
}

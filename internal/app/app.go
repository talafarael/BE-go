package app

import (
	"context"
	"errors"
	"fmt"
	"gin/internal/config"
	"gin/internal/controllers"
	"gin/internal/database/migrattion"
	"gin/internal/middleware"
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
	services := services.ConfigService(repo, &config)
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

func (app *App) configureRouter() {
	router := app.handler.Routing()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	repo := postgres.NewRepository(app.store.GetDB())
	// Regiuster service
	services := services.ConfigService(repo, app.config)
	middleware := middleware.ConfigMiddleware(repo, app.config)

	controller := controllers.NewBaseController(services, middleware)
	controller.RegisterRoutes(router)
}

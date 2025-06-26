package server

import (
	"context"
	"gin/internal/infrastructure/config"
	"gin/pkg/HttpServer"
	"gin/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer HttpServer.Server
	httpNet    http.Server
}

func (s *Server) HTTPServer(port string, router *gin.Engine) error {
	_, err := s.httpServer.HTTPServer(port, router)

	// s = &Server{httpServer: s.httpServer, httpNet: *httpServer}

	return err
}

func (s *Server) Run(config *config.Config, h *handler.Handler) error {
	// init server add addr and base_controller
	err := s.HTTPServer(config.BindAddr, h.Routing())

	return err
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpNet.Shutdown(ctx)
}

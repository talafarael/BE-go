package handler

import (
	"gin/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.Engine
	// services services.Service
}

func NewHandler(services services.Service) *Handler {
	return &Handler{
		router: gin.New(),
		// services: services,
	}
}

func (h *Handler) Routing() *gin.Engine {
	return h.router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

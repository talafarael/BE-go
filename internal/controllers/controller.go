package controllers

import (
	"gin/internal/services"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}
type BaseController struct {
	service     services.Service
	controllers []Controller
}

func NewBaseController(service services.Service) *BaseController {
	return &BaseController{
		service:     service,
		controllers: []Controller{},
	}
}

func (b *BaseController) AddSingleController(controllerFunc func(base *BaseController) Controller) {
	b.controllers = append(b.controllers, controllerFunc(b))
}

func (b *BaseController) RegisterRoutes(router *gin.Engine) {
	for _, controller := range b.controllers {
		controller.RegisterRoutes(router)
	}
}

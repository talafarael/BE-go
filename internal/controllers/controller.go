package controllers

import (
	"gin/internal/middleware"
	"gin/internal/services"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}
type BaseController struct {
	service *services.Service
	*AuthController
	*UserController
	*SwaggerController

	AuthMiddleware *middleware.AuthMiddleware
}

func NewBaseController(service *services.Service, authMiddleware *middleware.AuthMiddleware) *BaseController {
	return &BaseController{
		service:           service,
		AuthController:    NewAuthController(*service),
		UserController:    NewUserController(*service, authMiddleware),
		SwaggerController: NewSwaggerController(),
		AuthMiddleware:    authMiddleware,
	}
}

func (bc *BaseController) RegisterRoutes(router *gin.Engine) {
	bc.AuthController.AuthRoutes(router)
	bc.UserController.UserRoutes(router)
	bc.SwaggerController.SwaggerRoutes(router)
}

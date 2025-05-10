package controllers

import (
	"gin/internal/services"
	"gin/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}
type BaseController struct {
	service *services.Service
	*AuthController
	*UserController
	AuthMiddleware *middleware.AuthMiddleware
}

func NewBaseController(service *services.Service, authMiddleware *middleware.AuthMiddleware) *BaseController {
	return &BaseController{
		service:        service,
		AuthController: NewAuthController(*service),
		UserController: NewUserController(*service, authMiddleware),
		AuthMiddleware: authMiddleware,
	}
}

func (bc *BaseController) RegisterRoutes(router *gin.Engine) {
	bc.AuthController.AuthRoutes(router)
	bc.UserController.UserRoutes(router)
}

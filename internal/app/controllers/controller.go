package controllers

import (
	"gin/internal/app/services"
	auth_controller "gin/internal/auth/auth_controllers"
	user_controller "gin/internal/user/user_controllers"
	"gin/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}
type BaseController struct {
	service *services.Service
	*auth_controller.AuthController
	*user_controller.UserController
	AuthMiddleware *middleware.AuthMiddleware
}

func NewBaseController(service *services.Service, authMiddleware *middleware.AuthMiddleware) *BaseController {
	return &BaseController{
		service:        service,
		AuthController: auth_controller.NewAuthController(*service),
		UserController: user_controller.NewUserController(*service, authMiddleware),
		AuthMiddleware: authMiddleware,
	}
}

func (bc *BaseController) RegisterRoutes(router *gin.Engine) {
	bc.AuthController.AuthRoutes(router)
	bc.UserController.UserRoutes(router)
}

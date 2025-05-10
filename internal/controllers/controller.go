package controllers

import (
	"gin/internal/services"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}
type BaseController struct {
	service services.Service
	*AuthController
	*UserController
}

func NewBaseController(service services.Service) *BaseController {
	return &BaseController{
		service:        service,
		AuthController: NewAuthController(service),
		UserController: NewUserController(service),
	}
}

func (bc *BaseController) RegisterRoutes(router *gin.Engine) {
	bc.AuthController.AuthRoutes(router)
	bc.UserController.UserRoutes(router)
}

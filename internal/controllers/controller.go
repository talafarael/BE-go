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
	*VacancyController
	AuthMiddleware *middleware.AuthMiddleware
}

func NewBaseController(service *services.Service, authMiddleware *middleware.AuthMiddleware) *BaseController {
	return &BaseController{
		service:           service,
		AuthController:    NewAuthController(*service),
		UserController:    NewUserController(*service, authMiddleware),
		VacancyController: NewVacancyController(*service, authMiddleware),
		SwaggerController: NewSwaggerController(),
		AuthMiddleware:    authMiddleware,
	}
}

func (bc *BaseController) RegisterRoutes(router *gin.Engine) {
	bc.AuthController.AuthRoutes(router)
	bc.UserController.UserRoutes(router)
	bc.VacancyController.VacancyRoutes(router)
	bc.SwaggerController.SwaggerRoutes(router)
}

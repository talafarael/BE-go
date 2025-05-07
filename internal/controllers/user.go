package controllers

import (
	"gin/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
	userService services.UserService
}

func NewUserController(base *BaseController) Controller {
	return &UserController{
		userService: base.service,
	}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	getUser := uc.userService.Get("af")
	ctx.JSON(http.StatusCreated, getUser)
}

func (u *UserController) RegisterRoutes(router *gin.Engine) {
	router.GET("/user", u.GetUser)
}

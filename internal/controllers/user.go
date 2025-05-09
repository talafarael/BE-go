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

// GetUsers return list of all users from the database
// @Summary return list of all
// @Description return list of all users from the database
// @Tags Users
// @Success 200 {object} models.User
// @Router /user [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	getUser := uc.userService.Get("af")
	ctx.JSON(http.StatusCreated, getUser)
}

// RegisterRoutes sets up the routes for the UserController.
func (u *UserController) RegisterRoutes(router *gin.Engine) {
	router.GET("/user", u.GetUser)
}

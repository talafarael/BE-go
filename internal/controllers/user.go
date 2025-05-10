package controllers

import (
	"gin/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
	userService services.Service
}

func NewUserController(service services.Service) *UserController {
	return &UserController{
		userService: service,
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
func (u *UserController) UserRoutes(router *gin.Engine) {
	router.GET("/user", u.GetUser)
}

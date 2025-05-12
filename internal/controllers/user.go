package controllers

import (
	"gin/internal/middleware"
	"gin/internal/models"
	"gin/internal/services"
	"net/http"

	response_error "gin/pkg/error"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService    services.Service
	AuthMiddleware *middleware.AuthMiddleware
}

func NewUserController(service services.Service, authMiddleware *middleware.AuthMiddleware) *UserController {
	return &UserController{
		userService:    service,
		AuthMiddleware: authMiddleware,
	}
}

// GetUsers return  user from the database
// @Summary return user by token
// @Description return  user from the database by jwt token
// @Tags Users
// @Success 200 {object} userModels.UserResponse
// @Failure      404  {object}  response_error.ResponseError
// @Router /user [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	user, err := uc.userService.Get(ctx)
	if err != nil {
		response_error.HandlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, models.UserResponse{User: user})
}

// RegisterRoutes sets up the routes for the UserController.
func (u *UserController) UserRoutes(router *gin.Engine) {
	authMiddleware := *u.AuthMiddleware
	userGroup := router.Group("/user")
	userGroup.Use(authMiddleware.CheckUserMiddleware())
	userGroup.GET("/", u.GetUser)
}

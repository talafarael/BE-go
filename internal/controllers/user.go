package controllers

import (
	models "gin/internal/models/user"
	"gin/internal/services"
	"gin/pkg/middleware"
	"net/http"

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

// GetUsers return list of all users from the database
// @Summary return list of all
// @Description return list of all users from the database
// @Tags Users
// @Success 200 {object} models.User
// @Router /user [get]
func (uc *UserController) GetUser(ctx *gin.Context) {
	userValue, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := userValue.(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// RegisterRoutes sets up the routes for the UserController.
func (u *UserController) UserRoutes(router *gin.Engine) {
	authMiddleware := *u.AuthMiddleware
	userGroup := router.Group("/user")
	userGroup.Use(authMiddleware.CheckUserMiddleware())
	userGroup.GET("/", u.GetUser)
}

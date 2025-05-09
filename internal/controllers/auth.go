package controllers

import (
	"gin/internal/models"
	"gin/internal/services"
	"gin/pkg/handler_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
	authService services.AuthService
}

func (uc *AuthController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser, err := uc.authService.Register(&user)
	if err != nil {
		handler_errors.HandlerError(ctx, err)
	}
	ctx.JSON(http.StatusCreated, newUser)
}

func (u *AuthController) RegisterRoutes(router *gin.Engine) {
	router.POST("/auth/register", u.CreateUser)
}

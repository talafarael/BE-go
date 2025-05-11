package controllers

import (
	"gin/internal/dto"
	"gin/internal/models"
	"gin/internal/services"
	response_error "gin/pkg/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.Service
}

func NewAuthController(service services.Service) *AuthController {
	return &AuthController{
		service: service,
	}
}

// Register
// @Summary      Register account
// @Description  Register account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  userModels.RegisterDto  true  "User registration data"
// @Success      200  {object}  userModels.AuthReponse
// @Failure      404  {object}  response_error.ResponseError
// @Failure      500  {object}  response_error.ResponseError
// @Router       /auth/register [post]
func (uc *AuthController) Register(ctx *gin.Context) {
	var user dto.RegisterDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response_error.HandlerError(ctx, err)
		return
	}
	token, err := uc.service.Register(&user)
	if err != nil {
		response_error.HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, models.AuthReponse{
		Token: token,
	})
}

// Login
// @Summary      Login account
// @Description  Login account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  userModels.LoginDto  true  "User login data"
// @Success      200  {object}  userModels.AuthReponse
// @Failure      404  {object}  response_error.ResponseError
// @Failure      500  {object}  response_error.ResponseError
// @Router       /auth/login [post]
func (uc *AuthController) Login(ctx *gin.Context) {
	var user dto.LoginDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response_error.HandlerError(ctx, err)
		return
	}
	token, err := uc.service.Login(&user)
	if err != nil {
		response_error.HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, models.AuthReponse{
		Token: token,
	})
}

func (u *AuthController) AuthRoutes(router *gin.Engine) {
	router.POST("/auth/register", u.Register)
	router.POST("/auth/login", u.Login)
}

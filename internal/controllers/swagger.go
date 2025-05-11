package controllers

import (
	"gin/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type SwaggerController struct{}

func NewSwaggerController() *SwaggerController {
	return &SwaggerController{}
}

func (u *SwaggerController) SwaggerRoutes(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

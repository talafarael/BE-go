package controllers

import (
	"gin/internal/dto"
	"gin/internal/middleware"
	"gin/internal/services"
	response_error "gin/pkg/error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VacancyController struct {
	service        services.Service
	authMiddleware *middleware.AuthMiddleware
}

func NewVacancyController(service services.Service, authMiddleware *middleware.AuthMiddleware) *VacancyController {
	return &VacancyController{
		service:        service,
		authMiddleware: authMiddleware,
	}
}

func (vc *VacancyController) CreateVacancy(ctx *gin.Context) {
	var vacancy dto.CreateVacancyDto

	if err := ctx.ShouldBindJSON(&vacancy); err != nil {
		response_error.HandlerError(ctx, err)
		return
	}
	res, err := vc.service.CreateVacancy(ctx, &vacancy)
	if err != nil {
		response_error.HandlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (vc *VacancyController) UpdateVacancy(ctx *gin.Context) {
	var vacancy dto.UpdateVacancyDto
	id := ctx.Param("id")
	if id == "" {
		response_error.HandlerError(ctx, response_error.ErrIDEmpty)
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response_error.HandlerError(ctx, response_error.ErrIDFormtaion)
		return
	}
	if err := ctx.ShouldBindJSON(&vacancy); err != nil {
		response_error.HandlerError(ctx, err)
		return
	}
	res, err := vc.service.UpdateVacancy(ctx, &vacancy, uint(idUint))
	if err != nil {
		response_error.HandlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (vc *VacancyController) DeleteVacancity(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response_error.HandlerError(ctx, response_error.ErrIDEmpty)
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response_error.HandlerError(ctx, response_error.ErrIDFormtaion)
		return
	}
	res, err := vc.service.DeleteVacancy(ctx, uint(idUint))
	if err != nil {
		response_error.HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (vc *VacancyController) VacancyRoutes(router *gin.Engine) {
	authMiddleware := *vc.authMiddleware
	router.POST("/vacancy", authMiddleware.CheckUserMiddleware(), vc.CreateVacancy)
	router.PUT("/vacancy/:id", authMiddleware.CheckUserMiddleware(), vc.UpdateVacancy)
	router.DELETE("/vacancy/:id", authMiddleware.CheckUserMiddleware(), vc.DeleteVacancity)
}

package handler_errors

import (
	"errors"
	"gin/internal/domain/customErrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandlerError(ctx *gin.Context, customError error) {
	switch {
	case errors.Is(customError, customErrors.ErrUserAlreadyExists):
		ctx.JSON(http.StatusConflict, ErrorResponse{Error: "User already exists"})
	default:
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Internal server error"})
	}
}

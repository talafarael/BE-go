package response_error

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerError(ctx *gin.Context, err error) {
	var respErr *ResponseError
	if errors.As(err, &respErr) {
		ctx.JSON(respErr.Code, gin.H{"error": respErr.Message})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

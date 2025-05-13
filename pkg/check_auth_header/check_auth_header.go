package check_auth_header

import (
	response_error "gin/pkg/error"
	"strings"

	"github.com/gin-gonic/gin"
)

// CheckAuthHeader
type CheckAuthHeader struct{}

func (ca *CheckAuthHeader) ExtractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	const prefix = "Bearer "

	if strings.HasPrefix(authHeader, prefix) {
		return strings.TrimPrefix(authHeader, prefix), nil
	}

	return "", response_error.ErrForbidden
}

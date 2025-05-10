package middleware

import (
	"gin/internal/repository"
	check_auth_header "gin/pkg/checkAuthHeader"
	response_error "gin/pkg/error"
	"gin/pkg/jwt"
	"log"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	CheckUserMiddleware() gin.HandlerFunc
}

type authMiddleware struct {
	AuthMiddlewareOptions
}
type AuthMiddlewareOptions struct {
	CheckAuthHeader *check_auth_header.CheckAuthHeader
	JwtService      *jwt.JwtService
	Repo            *repository.Store
}

func NewAuthMiddleware(options *AuthMiddlewareOptions) AuthMiddleware {
	return &authMiddleware{
		AuthMiddlewareOptions: *options,
	}
}

func (a *authMiddleware) CheckUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		strToken, err := a.CheckAuthHeader.ExtractToken(c)
		if err != nil {
			response_error.HandlerError(c, err)
			c.Abort()
		}
		id, err := a.JwtService.VerifyToken(strToken)
		if err != nil {
			log.Printf("1")
			response_error.HandlerError(c, err)
			c.Abort()
			return
		}
		user, err := (*a.Repo).User().GetUserByID(id)
		if err != nil {

			response_error.HandlerError(c, err)
			c.Abort()
			return
		}
		c.Set("user", user)

		c.Next()
	}
}

package jwt_test

import (
	response_error "gin/pkg/error"
	jwtService "gin/pkg/jwt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJwtService(t *testing.T) {
	secretKey := "test_secret_key"
	service := jwtService.NewJwtService(secretKey)

	t.Run("Create and Verify Token", func(t *testing.T) {
		userID := uint(42)

		tokenStr, err := service.CreateToken(userID)
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenStr)

		verifiedID, err := service.VerifyToken(tokenStr)
		assert.NoError(t, err)
		assert.Equal(t, userID, verifiedID)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		_, err := service.VerifyToken("invalid.token.here")
		assert.ErrorIs(t, err, response_error.ErrInvalidCredentials)
	})

	t.Run("Expired Token", func(t *testing.T) {
		claims := jwt.MapClaims{
			"id":  uint(99),
			"exp": time.Now().Add(-time.Hour).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString([]byte(secretKey))
		assert.NoError(t, err)

		_, err = service.VerifyToken(tokenStr)
		assert.Error(t, err)
	})
}

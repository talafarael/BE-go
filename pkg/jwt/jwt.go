package jwt

import (
	response_error "gin/pkg/error"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey string
}
type JwtService interface {
	CreateToken(id uint) (string, error)
	VerifyToken(tokenString string) (uint, error)
}

func NewJwtService(secretKey string) JwtService {
	return &jwtService{
		secretKey: secretKey,
	}
}

func (j *jwtService) CreateToken(id uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", response_error.ErrJWTCreationFailed
	}
	return strToken, nil
}

func (j *jwtService) VerifyToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, response_error.ErrInvalidCredentials
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return 0, response_error.ErrInvalidCredentials
	}

	if !token.Valid {
		return 0, response_error.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, response_error.ErrInvalidClaims
	}

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, response_error.ErrInvalidIDClaim
	}

	return uint(idFloat), nil
}

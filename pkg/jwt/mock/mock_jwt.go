package mock_jwt

import (
	response_error "gin/pkg/error"
	"gin/pkg/jwt"
	"strconv"
)

type mockJwtService struct {
	secretKey string
}

type MockJwtService interface {
	CreateToken(id uint) (string, error)
	VerifyToken(tokenString string) (uint, error)
}

func NewMockJwtService(secret string) jwt.JwtService {
	return &mockJwtService{
		secretKey: secret,
	}
}

func (m *mockJwtService) CreateToken(id uint) (string, error) {
	token := strconv.FormatUint(uint64(id), 10) + "_jwt"
	return token, nil
}

func (m *mockJwtService) VerifyToken(tokenString string) (uint, error) {
	if tokenString == "1_jwt" {
		return 1, nil
	}
	return 0, response_error.ErrInvalidCredentials
}

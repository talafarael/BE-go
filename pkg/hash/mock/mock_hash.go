package mock_hash

import (
	"gin/pkg/hash"
)

type mockHashService struct{}

func NewHashService() hash.HashService {
	return &mockHashService{}
}

func (m *mockHashService) HashPassword(password string) (string, error) {
	return "hash_password", nil
}

func (m *mockHashService) CheckPasswordHash(password, hash string) bool {
	return password == "password"
}

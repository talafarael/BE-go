package services_test

import (
	"gin/internal/dto"
	mocks_repository "gin/internal/repository/mocks_repo"
	"gin/internal/services"
	mock_hash "gin/pkg/hash/mock"
	mock_jwt "gin/pkg/jwt/mock"
	test_case "gin/pkg/test"
	"testing"
)

func TestAuth(t *testing.T) {
	testsLogin := []test_case.TestCase[*dto.LoginDto, string]{
		{
			Name: "Existing jwt token",
			Data: &dto.LoginDto{
				Email:    "test@example.com",
				Password: "password",
			},
			Res:     "1_jwt",
			Err:     nil,
			WantErr: false,
		},
	}
	for _, ttl := range testsLogin {
		t.Run(ttl.Name, func(t *testing.T) {
			mockJwtService := mock_jwt.NewMockJwtService("secret")
			mockHashService := mock_hash.NewHashService()
			mockRepo := mocks_repository.NewRepository()
			authService := services.NewAuthService(services.AuthServiceOptions{
				Repo:        mockRepo,
				JwtService:  mockJwtService,
				HashService: mockHashService,
			})
			login, err := authService.Login(ttl.Data)
			if (err != nil) != ttl.WantErr {
				t.Errorf("AuthService.Login() error = %v, wantErr %v", err, ttl.WantErr)
			}
			if ttl.Res == "" && login != "" {
				t.Errorf("AuthService.Login() = %v, want nil", login)
			} else if login != ttl.Res {
				t.Errorf("AuthService.Login() = %v, want %v", login, ttl.Res)
			} else if ttl.Err != err && ttl.WantErr {
				t.Errorf("AuthService.Login() = %v, want err %v", err, ttl.Err)
			}
		})
	}
}

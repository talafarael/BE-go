package auth_services_test

import (
	"gin/internal/dto"
	mocks_repository "gin/internal/repository/mocks_repo"
	"gin/internal/services"
	response_error "gin/pkg/error"
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
		{
			Name: "Error password not correct",
			Data: &dto.LoginDto{
				Email:    "test@example.com",
				Password: "not_correct_password",
			},
			Res:     "",
			Err:     response_error.ErrUnauthorized,
			WantErr: true,
		},
		{
			Name: "Error password not correct",
			Data: &dto.LoginDto{
				Email:    "test1@example.com",
				Password: "password",
			},
			Res:     "",
			Err:     response_error.ErrUserNotFound,
			WantErr: true,
		},
	}
	t.Run("AuthService.Login", func(t *testing.T) {
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
					t.Error
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
	})
	testsRegister := []test_case.TestCase[*dto.RegisterDto, string]{
		{
			Name: "Existing jwt token",
			Data: &dto.RegisterDto{
				Email:    "test1@example.com",
				Password: "password",
				Name:     "fara",
			},
			Res:     "2_jwt",
			Err:     nil,
			WantErr: false,
		},
		{
			Name: "Existing jwt token",
			Data: &dto.RegisterDto{
				Email:    "test@example.com",
				Password: "password",
				Name:     "fara",
			},
			Res:     "",
			Err:     response_error.ErrUserAlredy,
			WantErr: true,
		},
	}
	t.Run("AuthService.Register", func(t *testing.T) {
		for _, ttl := range testsRegister {
			t.Run(ttl.Name, func(t *testing.T) {
				mockJwtService := mock_jwt.NewMockJwtService("secret")
				mockHashService := mock_hash.NewHashService()
				mockRepo := mocks_repository.NewRepository()
				authService := services.NewAuthService(services.AuthServiceOptions{
					Repo:        mockRepo,
					JwtService:  mockJwtService,
					HashService: mockHashService,
				})
				register, err := authService.Register(ttl.Data)

				if (err != nil) != ttl.WantErr {
					t.Errorf("AuthService.Register() error = %v, wantErr %v", err, ttl.WantErr)
				}
				if ttl.Res == "" && register != "" {
					t.Errorf("AuthService.Register() = %v, want nil", register)
				} else if register != ttl.Res {
					t.Errorf("AuthService.Register() = %v, want %v", register, ttl.Res)
				} else if ttl.Err != err && ttl.WantErr {
					t.Errorf("AuthService.Register() = %v, want err %v", err, ttl.Err)
				}
			})
		}
	})
}

package services_test

import (
	"gin/internal/models"
	mocks_repository "gin/internal/repository/mocks_repo"
	"gin/internal/services"
	"testing"

	response_error "gin/pkg/error"
	test_case "gin/pkg/test"

	"github.com/gin-gonic/gin"
)

func TestUser(t *testing.T) {
	ctx := &gin.Context{}
	user := &models.User{
		ID:       1,
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	ctx.Set("user", user)

	tests := []test_case.TestCase[*gin.Context, *models.User]{
		{
			Name: "Existing user",
			Data: ctx,
			Res: &models.User{
				ID:       1,
				Name:     "testuser",
				Email:    "test@example.com",
				Password: "password",
			},
			Err:     nil,
			WantErr: false,
		},
		{
			Name:    "Empty user",
			Data:    &gin.Context{},
			Res:     &models.User{},
			Err:     response_error.ErrUnauthorized,
			WantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockRepo := mocks_repository.NewRepository()

			gotUser, err := services.NewUserService(services.UserServiceOptions{Repo: mockRepo}).Get(tt.Data)

			if (err != nil) != tt.WantErr {
				t.Errorf("UserService.GetUser() error = %v, wantErr %v", err, tt.WantErr)
				return
			}

			if tt.Res == nil && gotUser != nil {
				t.Errorf("UserService.GetUser() = %v, want nil", gotUser)
			} else if tt.Res != nil {
				if gotUser == nil {
					t.Errorf("UserService.GetUser() = nil, want %v", tt.Res)
				} else if gotUser.ID != tt.Res.ID ||
					gotUser.Name != tt.Res.Name ||
					gotUser.Email != tt.Res.Email {
					t.Errorf("UserService.GetUser() = %v, want %v", gotUser, tt.Res)
				}
			}
		})
	}
}

package user_dto

import "gin/internal/user/user_models"

type UserResponse struct {
	User user_models.User `json:"user"`
}

package models

import "gin/internal/user/user_models"

type UserResponse struct {
	User user_models.User `json:"user"`
}

package dto

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

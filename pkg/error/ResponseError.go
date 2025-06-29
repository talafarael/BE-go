package response_error

import (
	"fmt"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponseError(code int, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("Code\t%d, message\t%s", e.Code, e.Message)
}

var (
	ErrInvalidCredentials = NewResponseError(401, "invalid credentials")
	ErrInternalServer     = NewResponseError(500, "Internal server error")
	// jwt
	ErrInvalidToken      = NewResponseError(401, "invalid token")
	ErrInvalidClaims     = NewResponseError(400, "could not parse claims")
	ErrInvalidIDClaim    = NewResponseError(400, "id claim missing or invalid")
	ErrJWTCreationFailed = NewResponseError(500, "JWT error: failed to create token")
	// user
	ErrUserAlredy                = NewResponseError(401, "Email alredy exists")
	ErrUserNotFound              = NewResponseError(401, "User notr found")
	ErrPasswordOrEmailNotCorrect = NewResponseError(401, "Invalid email or password")
	// authoriztion err
	ErrUnauthorized = NewResponseError(401, "Unauthorized")
	ErrForbidden    = NewResponseError(403, "Forbidden")
	// vacancy
	ErrVacancyCreateError = NewResponseError(400, "Error creating vacancy")
	ErrVacancyNotFound    = NewResponseError(404, "Vacancy not found")
)

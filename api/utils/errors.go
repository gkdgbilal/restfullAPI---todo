package utils

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type HttpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e HttpError) Error() string {
	return e.Message
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Status defaults to 500
	status := http.StatusInternalServerError

	// Check if it's an HttpError type
	if e, ok := err.(*HttpError); ok {
		status = e.Status
	}

	return c.Status(status).JSON(&HttpError{
		Status:  status,
		Message: err.Error(),
	})
}

func NewError(status int, message string) *HttpError {
	return &HttpError{
		Status:  status,
		Message: message,
	}
}

var (
	ErrUnauthorized            = NewError(http.StatusUnauthorized, "Unauthorized user.")
	ErrUsernameAlreadyTaken    = NewError(http.StatusBadRequest, "User already taken.")
	ErrEmailAlreadyTaken       = NewError(http.StatusBadRequest, "Email already taken.")
	ErrPasswordNotAcceptable   = NewError(http.StatusBadRequest, "Password not acceptable.")
	ErrInvalidPasswordUsername = NewError(http.StatusBadRequest, "User or password invalid.")
	ErrInvalidId               = NewError(http.StatusBadRequest, "Invalid ID.")
	ErrEmptyId                 = NewError(http.StatusBadRequest, "ID can't be empty")
	ErrInvalidRequest          = NewError(http.StatusBadRequest, "Invalid request.")
	ErrFailedSave              = NewError(http.StatusServiceUnavailable, "We couldn't save your request. Please try again!")
	ErrFailedRead              = NewError(http.StatusServiceUnavailable, "We couldn't read your request. Please try again!")
	ErrInvalidQueryParam       = NewError(http.StatusBadRequest, "Your requested query params are invalid. Please try again.")
)

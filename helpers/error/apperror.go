package error

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewBadRequestAppError(msg string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: msg}

}
func NewNotFoundAppError(msg string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: msg}

}
func NewInternalServerAppError(msg string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: msg}

}
func NewUnauthorizedAppError(msg string) *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: msg}

}
func (apperror *AppError) Getmessage() *AppError {
	return &AppError{Message: apperror.Message}
}

package service

import (
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/dto"
)

type UserService interface {
	RegisterUser(dto.User) (dto.User, *error.AppError)
}

package repository

import (
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/domain"
)

type UserRepository interface {
	RegisterUser(domain.User) (domain.User, *error.AppError)
}

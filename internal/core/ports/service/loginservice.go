package service

import (
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/dto"
)

type LoginService interface {
	FetchLoginDetails(dto.LoginDto) (*dto.Authtoken, *error.AppError)
}

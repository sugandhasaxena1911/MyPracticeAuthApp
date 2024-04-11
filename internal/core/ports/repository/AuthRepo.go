package repository

import (
	coreerror "github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/domain"
)

type AuthRepository interface {
	FetchLoginDetails(username string, password string) (domain.LoginDetails, *coreerror.AppError)
}

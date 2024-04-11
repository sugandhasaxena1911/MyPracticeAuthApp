package service

import (
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/dto"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/ports/repository"
)

type LoginCoreService struct {
	Authrepo repository.AuthRepository
}

func (logincoreservice LoginCoreService) FetchLoginDetails(logindto dto.LoginDto) (*dto.Authtoken, *error.AppError) {
	domainlogin, err := logincoreservice.Authrepo.FetchLoginDetails(logindto.Username, logindto.Password)
	if err != nil {
		return nil, err
	}
	token, err := domainlogin.GenerateToken()
	if err != nil {
		return nil, err
	}
	authtoken := dto.Authtoken{Token: *token}
	return &authtoken, err

}
func NewLoginCoreService(authrepo repository.AuthRepository) LoginCoreService {
	return LoginCoreService{authrepo}

}

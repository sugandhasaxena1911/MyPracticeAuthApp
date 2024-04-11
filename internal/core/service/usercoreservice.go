package service

import (
	"database/sql"
	"log"

	coreerror "github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/domain"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/dto"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/ports/repository"
)

type UserCoreService struct {
	Userrepo repository.UserRepository
}

func (usrcoreservice UserCoreService) RegisterUser(userdto dto.User) (dto.User, *coreerror.AppError) {
	log.Println("Inside usercore service resgister user")
	var custid sql.NullString
	if userdto.Customerid == "" {
		custid.String = ""
		custid.Valid = false
	} else {
		custid.String = userdto.Customerid
		custid.Valid = true
	}

	domainuser := domain.User{Username: userdto.Username, Password: userdto.Password, Role: userdto.Role, Customerid: custid}
	log.Println("converted dto to domain user ", domainuser)
	domainuser, err := usrcoreservice.Userrepo.RegisterUser(domainuser)
	log.Println("After calling repo db register user ", domainuser)
	if err != nil {
		log.Println("eror from user repo db register user", err.Message)
		return userdto, err
	}
	userdto = domainuser.TouserDto()
	log.Println("After converting domain to dto  ", userdto)

	return userdto, nil
}

func NewUserCoreService(usrrepo repository.UserRepository) UserCoreService {
	return UserCoreService{Userrepo: usrrepo}

}

package repositoryDB

import (
	"database/sql"
	"fmt"
	"log"

	coreerror "github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/logger"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryDB struct {
	dbclient *sql.DB
}

func encryptUserPassword(textpass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(textpass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), err
}
func (userrepodb UserRepositoryDB) RegisterUser(user domain.User) (domain.User, *coreerror.AppError) {
	log.Println("inside user repo db ")
	db := userrepodb.dbclient
	sqlst := "insert into users(username,password,role,customer_id) values(?,?,?,?)"
	encryptedpass, e := encryptUserPassword(user.Password)
	if e != nil {
		return user, coreerror.NewBadRequestAppError("Cannot encrypt password")
	}
	_, err := db.Exec(sqlst, &user.Username, encryptedpass, &user.Role, &user.Customerid)
	if err != nil {
		logger.Error(fmt.Sprintln("db error while insertion : ", err))
		return user, coreerror.NewInternalServerAppError("DB error : cannot insert User")
	}
	logger.Info("Insert is successful")
	return user, nil

}

func NewUserRespositoryDB(client *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{dbclient: client}

}

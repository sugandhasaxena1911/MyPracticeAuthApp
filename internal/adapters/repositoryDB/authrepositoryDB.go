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

type AuthRepositoryDB struct {
	DBclient *sql.DB
}

func (authrepodb AuthRepositoryDB) FetchLoginDetails(username string, password string) (domain.LoginDetails, *coreerror.AppError) {
	var logindet domain.LoginDetails
	err := authrepodb.verifydetails(username, password)
	if err != nil {
		logger.Info(fmt.Sprintf("Error while verifying details %s", err.Message))
		return logindet, err
	}

	sqlst := "select username , role, customer_id from users where username= ?"
	row := authrepodb.DBclient.QueryRow(sqlst, username)
	e := row.Scan(&logindet.Username, &logindet.Role, &logindet.Customerid)
	if e != nil {
		log.Println("Error while scanning the records ", e)
		return logindet, coreerror.NewUnauthorizedAppError("User not identified/Invalid username")
	}

	return logindet, nil

}

func (authrepodb AuthRepositoryDB) verifydetails(username string, password string) *coreerror.AppError {

	sqlst := "select password from users where username= ?"
	var pass string
	row := authrepodb.DBclient.QueryRow(sqlst, username)
	err := row.Scan(&pass)
	if err != nil {
		logger.Info(fmt.Sprintf("Error while fetching username %s", err))
		return coreerror.NewUnauthorizedAppError("Invalid username")
	}

	// check password match for username
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	if err != nil {
		logger.Info(fmt.Sprintf("Passwords not matched  %s", err))
		return coreerror.NewUnauthorizedAppError("Wrong password")
	}

	return nil

}

func NewAuthRepositoryDB(dbclient *sql.DB) AuthRepositoryDB {
	return AuthRepositoryDB{DBclient: dbclient}

}

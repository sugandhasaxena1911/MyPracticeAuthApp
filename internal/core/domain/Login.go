package domain

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/error"
)

const TOKENDURATION = time.Hour

type LoginDetails struct {
	Username   string
	Customerid sql.NullString
	Role       string
}

func (l LoginDetails) GenerateToken() (*string, *error.AppError) {
	//var token string
	var claims jwt.Claims

	if l.Role == "user" {
		claims = l.getclaimsforuser()

	}
	if l.Role == "admin" {
		claims = l.getclaimsforadmin()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	st := os.Getenv("SECRETKEY")
	tokenstring, e := token.SignedString([]byte(st))
	if e != nil {
		log.Println("e", e)
		return nil, error.NewUnauthorizedAppError("cannot generate Token")
	}

	return &tokenstring, nil

}

func (l LoginDetails) getclaimsforuser() jwt.MapClaims {
	return jwt.MapClaims{"username": l.Username,
		"customerid": l.Customerid, "role": l.Role, "exp": time.Now().Add(TOKENDURATION).Unix()}

}
func (l LoginDetails) getclaimsforadmin() jwt.MapClaims {
	return jwt.MapClaims{"username": l.Username,
		"role": l.Role, "exp": time.Now().Add(TOKENDURATION).Unix()}
}

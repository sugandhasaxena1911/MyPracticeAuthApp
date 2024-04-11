package domain

import (
	"database/sql"

	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/dto"
)

type User struct {
	Username   string
	Password   string
	Role       string
	Customerid sql.NullString
	CreatedOn  string
}

func (u User) TouserDto() dto.User {
	var custid string
	var createdon string

	if u.Customerid.Valid {
		custid = u.Customerid.String

	} else {
		custid = ""
	}
	/*if u.CreatedOn.Valid {
		createdon = u.CreatedOn.String

	} else {
		createdon = ""
	}
	*/
	createdon = u.CreatedOn
	return dto.User{Username: u.Username, Password: u.Password, Role: u.Role, Customerid: custid, CreatedOn: createdon}
}

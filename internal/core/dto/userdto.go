package dto

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Customerid string `json:"customerid,omitempty"`
	CreatedOn  string `json:"createdon,omitempty"`
}

package models

type PasswordLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
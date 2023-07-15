package models

import "errors"

var (
	ErrNotValidUser = errors.New("not valid user values")
)

type User struct {
	ID       int    `json:"id" db:"id" `
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" db:"Phone"`
	Username string `json:"username" db:"Username" binding:"required"`
	Password string `json:"password" db:"Password_hash" binding:"required"`
}

type Auth struct {
	Username string `json:"username"`
}

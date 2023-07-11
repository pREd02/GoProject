package repository

import (
	"GoProject/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
	GetUserId(id int) (models.User, error)
}

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT  INTO %s (name,username,organName,address,phone,password_hash,role,biInIin) values($1,$2,$3,$4) RETURNING  id", "users")
	row := u.db.QueryRow(query, user.Name, user.Username, user.Phone, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("repo: create user: %w", err)
	}
	return id, nil
}

func (u *User) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", "users")
	err := u.db.Get(&user, query, username, password)
	if err != nil {
		return models.User{}, fmt.Errorf("repo: get user: %w", err)
	}
	return user, nil
}

func (u *User) GetUserId(id int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", "users")
	err := u.db.Get(&user, query, id)
	if err != nil {
		return models.User{}, fmt.Errorf("repo: get user id: %w", err)
	}
	return user, err
}

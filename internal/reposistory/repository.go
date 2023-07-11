package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	Organization
	Projects
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Organization: NewOrgan(db),
		Projects:     NewProject(db),
		Users:        NewUser(db),
	}
}

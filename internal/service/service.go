package service

import (
	"GoProject/internal/reposistory"
)

type Service struct {
	Organization
	Projects
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Organization: NewOrgan(repos.Organization),
		Projects:     NewProject(repos.Projects),
		Users:        NewUser(repos.Users),
	}
}

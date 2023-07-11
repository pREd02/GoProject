package service

import (
	"GoProject/internal/models"
	repository "GoProject/internal/reposistory"
)

type Organization interface {
	Create(organization models.Organization) (int, error)
	GetOrganById(id int) (models.Organization, error)
}

type OrganService struct {
	repo repository.Organization
}

func (o *OrganService) Create(organization models.Organization) (int, error) {
	return o.repo.Create(organization)
}

func (o *OrganService) GetOrganById(id int) (models.Organization, error) {
	return o.repo.GetOrganizationById(id)
}

func NewOrgan(repo repository.Organization) *OrganService {
	return &OrganService{repo: repo}
}

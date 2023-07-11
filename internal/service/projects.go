package service

import (
	"GoProject/internal/models"
	repository "GoProject/internal/reposistory"
)

type Projects interface {
	CreateProject(project models.Projects) error
	GetProjectByID(id int) (models.Projects, error)
	GetUsersByProjectID(projectID int) ([]models.User, error)
	SubmitApplication(projectID, userID int) error
}

type ProjectService struct {
	repo repository.Projects
}

func NewProject(repo repository.Projects) *ProjectService {
	return &ProjectService{repo: repo}
}
func (p *ProjectService) CreateProject(projects models.Projects) error {
	return p.repo.CreateProject(projects)
}
func (p *ProjectService) GetProjectByID(id int) (models.Projects, error) {
	return p.repo.GetProjectByID(id)
}
func (p *ProjectService) GetUsersByProjectID(projectID int) ([]models.User, error) {
	return p.repo.GetUsersByProjectID(projectID)
}
func (p *ProjectService) SubmitApplication(projectID, userID int) error {
	return p.repo.SubmitApplication(projectID, userID)
}

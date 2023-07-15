package repository

import (
	"GoProject/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Projects interface {
	CreateProject(project models.Projects) error
	GetProjectByID(id int) (models.Projects, error)
	GetUsersByProjectID(projectID int) ([]models.User, error)
	SubmitApplication(projectID, userID int) error
}

type Project struct {
	db *sqlx.DB
}

func NewProject(db *sqlx.DB) *Project {
	return &Project{db: db}
}
func (p *Project) CreateProject(project models.Projects) error {
	query := fmt.Sprintf("INSERT INTO %s ( name, organization_id, date, owner) VALUES (:name, :organization_id, :date, :owner)", "projects")
	_, err := p.db.Exec(query, project.Name, project.OrganizationID, project.Date, project.Owner)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) GetProjectByID(id int) (models.Projects, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", "projects")
	var project models.Projects
	err := p.db.QueryRow(query, id).Scan(&project.ID, &project.Name, &project.OrganizationID, &project.Date, &project.Owner)
	if err != nil {
		return project, err
	}

	return project, nil
}

func (p *Project) GetUsersByProjectID(projectID int) ([]models.User, error) {
	query := fmt.Sprintf("SELECT  users.* FROM %s JOIN %s ON users.id = project_users.user_id WHERE project_users.project_id = ?", "users", "project_users")
	rows, err := p.db.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *Project) SubmitApplication(projectID, userID int) error {
	query := fmt.Sprintf("INSERT INTO $s (project_id, user_id) VALUES (:project_id, :user_id)", "project_users")
	_, err := p.db.Exec(query, projectID, userID)
	if err != nil {
		return err
	}

	return nil
}

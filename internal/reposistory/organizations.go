package repository

import (
	"GoProject/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Organization interface {
	Create(organization models.Organization) (int, error)
	GetOrganizationById(id int) (models.Organization, error)
}

type Organ struct {
	db *sqlx.DB
}

func NewOrgan(db *sqlx.DB) *Organ {
	return &Organ{db: db}
}

func (o *Organ) Create(organ models.Organization) (int, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (name, bin, address, workers) VALUES ($1,$2,$3,$4)", "organizations")
	result, err := o.db.Exec(query, organ.Name, organ.Bin, organ.Address, organ.Workers)
	if err != nil {
		return 0, fmt.Errorf("repo: create products: %w", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("repo: create products: %w", err)
	}
	return int(id), nil
}

func (o *Organ) GetOrganizationById(id int) (models.Organization, error) {
	var organization models.Organization
	query := fmt.Sprintf("select * from %s where id = $1", "organizations")
	err := o.db.Get(&organization, query, id)
	if err != nil {
		return organization, fmt.Errorf("repo: get product by id: %w", err)
	}
	return organization, nil
}

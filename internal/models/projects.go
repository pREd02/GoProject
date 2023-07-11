package models

type Projects struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Date           string `json:"date" db:"date"`
	Owner          string `json:"owner" db:"owner"`
	OrganizationID int    `db:"organization_id"`
}

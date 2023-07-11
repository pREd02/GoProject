package models

type Organization struct {
	ID      int    `json:"id" db:"id"`
	Bin     string `json:"bin" db:"bin"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
	Workers string `json:"workers" db:"workers"`
}

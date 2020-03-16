package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Developer struct
type Developer struct {
	BaseModel
	Name       string         `json:"name,omitempty"`
	HeadOffice postgres.Jsonb `json:"head_office,omitempty"`
	Apartment  []Apartment    `json:"-" gorm:"foreignkey:DeveloperID"`
}

// TableName Developer
func (Developer) TableName() string {
	return "developer"
}

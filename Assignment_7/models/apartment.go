package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Apartment struct
type Apartment struct {
	BaseModel
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	DeveloperID int            `json:"-"`
	Developer   Developer      `json:"developer,omitempty"`
	Location    postgres.Jsonb `json:"location,omitempty"`
	Units       []Unit         `json:"sell_units,omitempty" gorm:"foreignkey:ApartmentID"`
}

// TableName Apartment
func (Apartment) TableName() string {
	return "apartment"
}

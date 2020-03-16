package models

import (
	"errors"
)

// Unit struct
type Unit struct {
	BaseModel
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Total       int     `json:"total,omitempty"`
	Available   int     `json:"available,omitempty"`
	Status      string  `json:"status,omitempty"`
	ApartmentID int     `json:"-"`
}

// TableName Unit
func (Unit) TableName() string {
	return "unit"
}

func (u *Unit) BeforeCreate() (err error) {
	if u.Status != "rent" && u.Status != "sell" {
		err = errors.New("Status value is only rent or sell")
	}

	if u.Total <= 0 {
		err = errors.New("Total must be higher than 0")
	}

	if u.Available > u.Total {
		err = errors.New("Total must be bigger than or equal of Available")
	}

	if u.Available == 0 {
		u.Available = u.Total
	}

	return
}

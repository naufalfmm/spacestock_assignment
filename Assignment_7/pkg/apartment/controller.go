package apartment

import "github.com/naufalfmm/spacestock_assignment/Assignment_7/models"

//Controller ...
type Controller interface {
	GetByID(id int) (models.Apartment, error)
	GetAll(filter map[string]string, order map[string]string, limit int, fromDate string, toDate string) ([]models.Apartment, error)
	CreateNewApartment(name string, description string, developerID int, location map[string]interface{}, units []models.Unit) error
	EditApartment(id int, updatedData map[string]interface{}) error
	DeleteApartment(id int) error
}

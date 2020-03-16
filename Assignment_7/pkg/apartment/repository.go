package apartment

import "github.com/naufalfmm/spacestock_assignment/Assignment_7/models"

//Repository of event face recognition interface
type Repository interface {
	GetByID(id int) (models.Apartment, error)
	GetAll(filter map[string]string, order map[string]string, limit int, fromDate string, toDate string) ([]models.Apartment, error)
	CreateApartment(name string, description string, developer models.Developer, location map[string]interface{}, units []models.Unit) (models.Apartment, error)
	EditApartment(id int, updatedData map[string]interface{}) error
	DeleteApartment(id int) error
}

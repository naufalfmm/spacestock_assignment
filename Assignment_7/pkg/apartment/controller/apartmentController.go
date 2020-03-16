package controller

import (
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/models"
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/pkg/apartment"
)

type apartmentController struct {
	apartmentRepository apartment.Repository
}

func NewApartmentController(apartmentRepo apartment.Repository) apartment.Controller {
	return &apartmentController{
		apartmentRepository: apartmentRepo,
	}
}

func (ac *apartmentController) GetByID(id int) (models.Apartment, error) {
	var data models.Apartment

	data, err := ac.apartmentRepository.GetByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (ac *apartmentController) GetAll(filter map[string]string, order map[string]string, limit int, fromDate string, toDate string) ([]models.Apartment, error) {
	var data []models.Apartment

	data, err := ac.apartmentRepository.GetAll(filter, order, limit, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ac *apartmentController) CreateNewApartment(name string, description string, developerID int, location map[string]interface{}, units []models.Unit) error {
	var developer models.Developer

	baseModelDeveloper := models.BaseModel{
		ID: developerID,
	}

	developer = models.Developer{
		BaseModel: baseModelDeveloper,
	}

	_, err := ac.apartmentRepository.CreateApartment(name, description, developer, location, units)
	if err != nil {
		return err
	}

	return nil
}

func (ac *apartmentController) EditApartment(id int, updatedData map[string]interface{}) error {
	err := ac.apartmentRepository.EditApartment(id, updatedData)
	if err != nil {
		return err
	}

	return nil
}

func (ac *apartmentController) DeleteApartment(id int) error {
	err := ac.apartmentRepository.DeleteApartment(id)
	if err != nil {
		return err
	}

	return nil
}

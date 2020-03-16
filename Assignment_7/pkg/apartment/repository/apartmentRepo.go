package repository

import (
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/models"
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/pkg/apartment"
)

type apartmentRepository struct {
	Conn *gorm.DB
}

//NewPostgresqlApartmentRepository create new object of apartment repository
func NewPostgresqlApartmentRepository(Connection *gorm.DB) apartment.Repository {
	return &apartmentRepository{Conn: Connection}
}

// GetByID --> get by id
func (ar *apartmentRepository) GetByID(id int) (models.Apartment, error) {
	var data models.Apartment

	fetchDb := ar.Conn.Where("id = ?", id).Preload("Units").Preload("Developer").First(&data)
	if fetchDb.Error != nil {
		return data, fetchDb.Error
	}

	return data, nil
}

func createWhereQuery(filter map[string]string, fromDate string, toDate string) (string, error) {
	whereQuery := ""

	for field, value := range filter {
		if whereQuery != "" {
			whereQuery += " AND "
		}

		if field == "created_at" {
			err := errors.New("created_at not accepted")

			return "", err
		} else {
			whereQuery = whereQuery + field + " LIKE '%" + value + "%'"
		}
	}

	if fromDate != "" || toDate != "" {
		if whereQuery != "" {
			whereQuery += " AND "
		}
	}

	if fromDate != "" && toDate != "" {
		whereQuery = whereQuery + "DATE(created_at) BETWEEN '" + fromDate + "' AND '" + toDate + "'"
	} else if fromDate != "" {
		whereQuery = whereQuery + "DATE(created_at) >= '" + fromDate + "'"
	} else if toDate != "" {
		whereQuery = whereQuery + "DATE(created_at) <= '" + toDate + "'"
	}

	return whereQuery, nil
}

func (ar *apartmentRepository) GetAll(filter map[string]string, order map[string]string, limit int, fromDate string, toDate string) ([]models.Apartment, error) {
	var data []models.Apartment

	whereQuery, err := createWhereQuery(filter, fromDate, toDate)
	if err != nil {
		return data, err
	}

	fetchDb := ar.Conn.Where(whereQuery)

	if order != nil {
		fetchDb = fetchDb.Order(order["field"] + " " + order["type"])
	}

	if limit != 0 {
		fetchDb = fetchDb.Limit(limit)
	}

	fetchDb = fetchDb.Preload("Units").Preload("Developer").Find(&data)
	if fetchDb.Error != nil {
		return nil, fetchDb.Error
	}

	return data, nil
}

func (ar *apartmentRepository) CreateApartment(name string, description string, developer models.Developer, location map[string]interface{}, units []models.Unit) (models.Apartment, error) {
	var data models.Apartment

	loc, _ := json.Marshal(location)
	locationJsonb := json.RawMessage(loc)

	data = models.Apartment{
		Name:        name,
		Developer:   developer,
		Description: description,
		Location:    postgres.Jsonb{locationJsonb},
		Units:       units,
	}

	fetchDb := ar.Conn.Create(&data)
	if fetchDb.Error != nil {
		return data, fetchDb.Error
	}

	fetchDb = ar.Conn.Save(&data)
	if fetchDb.Error != nil {
		return data, fetchDb.Error
	}

	return data, nil
}

func (ar *apartmentRepository) EditApartment(id int, updatedData map[string]interface{}) error {
	var data models.Apartment

	fetchDb := ar.Conn.Model(&data).Where("id = ?", id).Updates(updatedData)
	if fetchDb.Error != nil {
		return fetchDb.Error
	}

	return nil
}

func (ar *apartmentRepository) DeleteApartment(id int) error {
	var data models.Apartment

	fetchDb := ar.Conn.Where("id = ?", id).Delete(&data)
	if fetchDb.Error != nil {
		return fetchDb.Error
	}

	return nil
}

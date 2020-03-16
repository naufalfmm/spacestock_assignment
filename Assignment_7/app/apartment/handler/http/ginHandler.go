package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/models"
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/pkg/apartment"
)

type ApartmentPostData struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	DeveloperID int                    `json:"developer_id"`
	Location    map[string]interface{} `json:"location"`
	Units       []models.Unit          `json:"units"`
}

type GinHandler struct {
	apartmentController apartment.Controller
}

const baseApartmentPath = "/apartment"

func NewGinHandler(e *gin.Engine, apartmentControl apartment.Controller) {
	handler := &GinHandler{
		apartmentController: apartmentControl,
	}

	router := e.Group(baseApartmentPath)
	{
		router.POST("/", handler.CreateApartment)
		router.GET("/:id", handler.GetApartmentByID)
		router.GET("/", handler.GetApartments)
		router.PUT("/:id", handler.EditApartment)
		router.DELETE("/:id", handler.DeleteApartment)
	}
}

func (gh *GinHandler) DeleteApartment(c *gin.Context) {
	ID := c.Param("id")

	apartmentID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	err = gh.apartmentController.DeleteApartment(apartmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Success",
	})
}

func (gh *GinHandler) EditApartment(c *gin.Context) {
	var updatedData map[string]interface{}

	ID := c.Param("id")

	apartmentID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	if err := c.BindJSON(&updatedData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	err = gh.apartmentController.EditApartment(apartmentID, updatedData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok":      true,
		"message": "Success",
	})
}

func (gh *GinHandler) GetApartments(c *gin.Context) {
	order := c.QueryMap("order")
	if len(order) == 0 {
		order = nil
	}

	filter := c.QueryMap("filter")
	if len(filter) == 0 {
		filter = nil
	}

	fromDate := c.Query("from")
	toDate := c.Query("to")

	limitQuery := c.Query("limit")
	if limitQuery == "" {
		limitQuery = "0"
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	data, err := gh.apartmentController.GetAll(filter, order, limit, fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": data,
	})
}

func (gh *GinHandler) GetApartmentByID(c *gin.Context) {
	ID := c.Param("id")

	apartmentID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	data, err := gh.apartmentController.GetByID(int(apartmentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": data,
	})
}

func (gh *GinHandler) CreateApartment(c *gin.Context) {
	var data ApartmentPostData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	err := gh.apartmentController.CreateNewApartment(data.Name, data.Description, data.DeveloperID, data.Location, data.Units)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok":      true,
		"message": "Success",
	})
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	httpDeliveryApartment "github.com/naufalfmm/spacestock_assignment/Assignment_7/app/apartment/handler/http"
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/config"
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/models"

	apartmentController "github.com/naufalfmm/spacestock_assignment/Assignment_7/pkg/apartment/controller"

	apartmentRepository "github.com/naufalfmm/spacestock_assignment/Assignment_7/pkg/apartment/repository"
)

var db *gorm.DB
var conf = config.Setting()

//CORSMiddleware for setup web server cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupDB() string {
	host := conf.DBHost
	port := conf.DBPort
	dbname := conf.DBName
	username := conf.DBUsername
	password := conf.DBPass
	url := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbname + " password=" + password + " connect_timeout=10"

	return url
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	db, err := gorm.Open("postgres", setupDB())
	if err != nil {
		panic("Failed to connect database")
	}

	db.LogMode(true)

	apartmentRepo := apartmentRepository.NewPostgresqlApartmentRepository(db)
	apartmentControl := apartmentController.NewApartmentController(apartmentRepo)

	httpDeliveryApartment.NewGinHandler(r, apartmentControl)

	return r
}

func main() {
	port := conf.Port

	models.DBInit()

	r := setupRouter()
	r.Run(":" + port)

	defer models.Db.Close()
}

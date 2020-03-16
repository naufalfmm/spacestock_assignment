package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/naufalfmm/spacestock_assignment/Assignment_7/config"
)

var conf = config.Setting()

// Db Gorm
var Db *gorm.DB

// DBInit --> function of init DB
func DBInit() {
	host := conf.DBHost
	port := conf.DBPort
	dbname := conf.DBName
	username := conf.DBUsername
	password := conf.DBPass

	url := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbname + " password=" + password + " connect_timeout=10"
	var err error
	Db, err = gorm.Open("postgres", url)
	if err != nil {
		fmt.Println(err.Error())
	}
}

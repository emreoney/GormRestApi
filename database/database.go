package database

import (
	"fmt"
	"gomod/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const dbInformation = "host=localhost port=5432 user=postgres password=Alexabi12312 dbname=formProject sslmode=disable"

func Init() {
	DB, err = gorm.Open(postgres.Open(dbInformation), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	DB.AutoMigrate(models.User{})
}

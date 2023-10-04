package database

import (
	"log"

	"github.com/vikaputri/CRUD_Users_Tasks/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := "host=localhost user=vika password=password dbname=otto port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connection Database : ", err)
	}

	db.Debug().AutoMigrate(models.User{})
	db.AutoMigrate(&models.Task{})

}

func GetDB() *gorm.DB {
	return db
}

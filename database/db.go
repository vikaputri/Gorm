package database

import (
	"belajargolang/Hacktiv8/Sesi7/gorm1/models"
	"fmt"
	"log"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "p4ssw0rd"
	dbPort   = "5432"
	dbname   = "learning_gorm"
	db       *gorm.db
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s, dbname=%s sslmode=disable",
		host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting database:", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}

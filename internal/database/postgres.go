package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DNS = "host=postgresql user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func DB() *gorm.DB {
	if database == nil {
		dsn := "host=localhost user=postgres password=postgres dbname=web_marketplace port=5432 sslmode=disable TimeZone=Europe/Oslo"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			panic("Could not connect to database")
		}
		database = db
	}
	return database
}

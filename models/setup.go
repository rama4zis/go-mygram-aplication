package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "root"
	dbname   = "mygram"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// make auto migrate
	database.AutoMigrate(&User{})
	database.AutoMigrate(&SocialMedia{})
	database.AutoMigrate(&Photo{})
	database.AutoMigrate(&Comment{})

	DB = database
}

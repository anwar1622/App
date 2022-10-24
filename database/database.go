package database

import (
	"App/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var (
	db  *gorm.DB
	err error
)

func StartDB(config *Config) (*gorm.DB, error) {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)
	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	return db, err
}

func GetDB() *gorm.DB {
	return db
}

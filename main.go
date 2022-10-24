package main

import (
	"App/database"
	"App/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &database.Config{
		Host:     os.Getenv("PGHOST"),
		Port:     os.Getenv("PGPORT"),
		User:     os.Getenv("PGUSER"),
		Password: os.Getenv("PGPASSWORD"),
		DBName:   os.Getenv("PGDATABASE"),
		//Host:     os.Getenv("DB_HOST"),
		//Port:     os.Getenv("DB_PORT"),
		//User:     os.Getenv("DB_USER"),
		//Password: os.Getenv("DB_PASS"),
		//DBName:   os.Getenv("DB_NAME"),
	}
	_, err = database.StartDB(config)
	if err != nil {
		log.Fatal(err)
	}
	var PORT = os.Getenv("PORT")
	router.Routers().Run(":" + PORT)
}

package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Tito-74/Savannah-e-shop/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func DatabaseInit(){

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable client_encoding=UTF8 TimeZone=Africa/Nairobi", host, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

  db.AutoMigrate(&models.Customer{}, &models.Orders{})

  Database = DbInstance{Db: db}
}
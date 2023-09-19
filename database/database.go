package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Tito-74/Savannah-e-shop/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func DatabaseConnect(){

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	name := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")

	fmt.Println("name: ", name)

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable client_encoding=UTF8 TimeZone=Africa/Nairobi", host, user, pass, name, port)
	// fmt.Sprintf("postgres://%v:%v@%v:%v/%v", user, pass,host,port,name)
	fmt.Println("dbURL: ", dbURL)
	// fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable client_encoding=UTF8 TimeZone=Africa/Nairobi", host, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})


    if err != nil {
        log.Fatalln(err)
    }
	db.Logger = logger.Default.LogMode(logger.Info)
  db.AutoMigrate(&models.Customer{}, &models.Orders{})

   Database = DbInstance{Db: db}
}
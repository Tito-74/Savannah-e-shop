package main

import (
	"log"
	"net/http"

	"github.com/Tito-74/Savannah-e-shop/controller"
	"github.com/Tito-74/Savannah-e-shop/database"
	"github.com/gorilla/mux"
)




func main() {
	database.DatabaseInit()
	router := mux.NewRouter()

	router.HandleFunc("/", controller.HelloWorld).Methods(http.MethodGet)
	router.HandleFunc("/customer",  controller.CreateCustomerDetails).Methods(http.MethodPost)
	router.HandleFunc("/order", controller.CreateOrderDetails).Methods(http.MethodPost)

	log.Println("API is running!")
	log.Fatal(http.ListenAndServe(":3000", router))

}
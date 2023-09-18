package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Tito-74/Savannah-e-shop/models"
	"github.com/edwinwalela/africastalking-go/pkg/sms"
	"github.com/joho/godotenv"
)

// send message function
func SendMessage(name string, phone string, order *models.Orders) *sms.Response {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	apiKey := os.Getenv("AFRICASTALKING_APIKEY")
	userName := os.Getenv("AFRICASTALKING_USERNAME")
	sender := os.Getenv("AFRICASTALKING_SHORTCODE")
	// Define Africa's Talking SMS client
	client := &sms.Client{
		ApiKey:    apiKey,
		Username:  userName,
		IsSandbox: true,
	}

	name = strings.Split(name, " ")[0]

	bulkRequest := &sms.BulkRequest{
		To:            []string{phone},
		Message:       fmt.Sprintf("Hello %v, Your order of %v @ %v has been received", name, order.Item, order.Amount),
		From:          sender,
		BulkSMSMode:   true,
		RetryDuration: time.Hour,
	}

	response, err := client.SendBulk(bulkRequest)
	if err != nil {

		panic(err)
	}
	fmt.Println(response.Message)
	return &response
}
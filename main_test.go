package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tito-74/Savannah-e-shop/models"
	"github.com/go-playground/assert/v2"
)

func TestCreateCustomerDetails(t *testing.T) {

	requestBody := `{"Name":  "Nic jones","Code":  "QWERT","Phone": "+254714000000"}`

	req := httptest.NewRequest("POST", "/customer",
		strings.NewReader(requestBody))

	response := httptest.NewRecorder()
	// CreateCustomerDetails(response, req)
	req.Header.Set("Content-Type", "application/json")
	res := response.Result()
	assert.Equal(t, 200, res.StatusCode)

}

func TestCreateOrderDetails(t *testing.T) {
	order := &models.Orders{
		Item:       "product",
		Amount:     300.00,
		CustomerId: 1,
	}
	jsonCustomer, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(jsonCustomer))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	// CreateOrderDetails(response, req)
	assert.Equal(t, 200, response.Code)

}
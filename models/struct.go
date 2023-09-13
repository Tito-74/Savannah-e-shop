package models

import "time"

type Customer struct {
	Id        int `json:"customer_id" gorm:"primarykey"`
	CreatedAt time.Time
	Name      string `json:"name" gorm:"not null"`
	Code      string `json:"code" gorm:"not null"`
	Phone     string `json:"phone" gorm:"not null"`
}

type Orders struct {
	Id        int `json:"order_id" gorm:"primarykey"`
	CreatedAt time.Time
	Item      string  `json:"item" gorm:"not null"`
	Amount    float64 `json:"amount" gorm:"not null"`
	CustomerId int  `json:"customer_id" gorm:"not null"`
	Customer   Customer `gorm:"foreignkey:CustomerId"`
}

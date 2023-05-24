package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID           string    `gorm:"column:id"`
	Status       string `gorm:"column:status"`
	Items        string `gorm:"column:items"`
	Total        int    `gorm:"column:total"`
	CurrencyUnit string `gorm:"column:currency_unit"`
}

type OrderRequest struct {
	Status       string      `json:"status"`
	Items        []OrderItem `json:"items"`
	Total        int         `json:"total"`
	CurrencyUnit string      `json:"currency_unit"`
}

type OrderItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

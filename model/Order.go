package model

import (
	"gorm.io/gorm"
)

type Order struct {
	// OrderId string `gorm:primaryKey json:"order_id"`
	gorm.Model
	CustomerName string `json:"customer_name"`
}

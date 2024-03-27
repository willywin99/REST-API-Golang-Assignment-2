package repository

import (
	"assignment_2/model"

	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) Create(newOrder model.Order) (model.Order, error) {
	tx := or.db.Create(&newOrder)

	return newOrder, tx.Error
}

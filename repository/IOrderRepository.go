package repository

import "assignment_2/model"

type IOrderRepository interface {
	Create(newOrder model.Order) (model.Order, error)
}

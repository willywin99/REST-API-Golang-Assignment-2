package main

import (
	"assignment_2/controller"
	"assignment_2/lib"
	"assignment_2/model"
	"assignment_2/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := lib.InitDatabase()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Order{})
	if err != nil {
		panic(err)
	}

	orderRepository := repository.NewOrderRepository(db)
	orderController := controller.NewOrderController(orderRepository)

	ginEngine := gin.Default()

	ginEngine.POST("/order", orderController.Create)

	err = ginEngine.Run("localhost:8083")
	if err != nil {
		panic(err)
	}
}

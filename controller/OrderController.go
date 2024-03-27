package controller

import (
	"assignment_2/model"
	"assignment_2/repository"
	"assignment_2/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderRepository repository.IOrderRepository
}

func NewOrderController(orderRepository repository.IOrderRepository) *orderController {
	return &orderController{
		orderRepository: orderRepository,
	}
}

func (oc *orderController) Create(ctx *gin.Context) {
	var newOrder model.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	createdOrder, err := oc.orderRepository.Create(newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createdOrder, ""))
}

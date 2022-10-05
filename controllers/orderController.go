package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "task-sesi8/dto"
	"task-sesi8/helpers"
	"task-sesi8/models"
	"task-sesi8/repository"
)

type orderController struct {
	repo repository.OrderRepository
}

type Response struct {
	Message string `example:"Success"`
	Status  int    `example:"200"`
	Data    interface{}
}

func NewOrderController() *orderController {
	return &orderController{
		repository.NewOrderRepository(),
	}
}

// CreateOrder godoc
// @Summary Create order
// @Description Create order by product id and user id
// @Tags order
// @Accept json
// @Produce json
// @Param order body dto.InputOrder true "Create Order"
// @Success 201 {object} dto.ResponseOrder[dto.Order]
// @Router /orders [post]
func (oc *orderController) CreateOrder(c *gin.Context) {
	repo := oc.repo
	var req models.Order
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.AbortWithStatusJSON(400, "bad request")
		return
	}

	res, err := repo.CreateOrder(req)

	if err != nil {
		c.AbortWithStatusJSON(400, "bad request")
		return
	}

	payload := map[string]interface{}{
		"customerName": res.CustomerName,
		"id":           res.ID,
		"items":        res.Items,
	}
	response := helpers.APIResponse("Success create order", http.StatusCreated, "success", payload)
	c.JSON(http.StatusCreated, response)
}

// getOrders godoc
// @Summary list orders
// @Schemes
// @Description get all orders
// @Tags order
// @Accept json
// @Produce json
// @Success 200 {object} dto.ResponseOrder[[]dto.Order]
// @Router /orders [get]
func (oc *orderController) GetOrders(c *gin.Context) {
	repo := oc.repo
	res, err := repo.GetOrders()

	if err != nil {
		c.AbortWithStatusJSON(400, "bad request")
		return
	}
	response := helpers.APIResponse("Success get orders", http.StatusOK, "success", res)
	c.JSON(http.StatusCreated, response)
}

// UpdateOrder godoc
// @Summary Update order
// @Description Update order by orderId
// @Tags order
// @Accept json
// @Produce json
// @Param order body dto.InputOrder true "Update Order"
// @Param orderId path string true "Id from order"
// @Success 200 {object} dto.ResponseOrder[dto.Order]
// @Router /orders/{orderId} [put]
func (oc *orderController) UpdateOrder(c *gin.Context) {
	repo := oc.repo
	var req models.Order
	err := c.ShouldBindJSON(&req)
	id := c.Param("id")

	fmt.Println("iddd", id)

	if err != nil {
		c.AbortWithStatusJSON(400, "bad request")
		return
	}

	res, err := repo.UpdateOrder(id, req)

	if err != nil {
		c.AbortWithStatusJSON(400, "bad request")
		return
	}
	response := helpers.APIResponse("Success update order", http.StatusOK, "success", res)
	c.JSON(http.StatusCreated, response)
}

// DeleteOrder godoc
// @Summary Delete order
// @Description Delete order by orderId
// @Tags order
// @Accept json
// @Produce json
// @Param orderId path string true "Id from order"
// @Success 200 {object} Response
// @Router /orders/{orderId} [delete]
func (oc *orderController) DeleteOrder(c *gin.Context) {
	repo := oc.repo
	id := c.Param("id")
	res, err := repo.DeleteOrder(id)

	if err != nil {
		c.AbortWithStatusJSON(400, "bad request")
		return
	}
	response := helpers.APIResponse("Success delete order", http.StatusOK, "success", res)
	c.JSON(http.StatusOK, response)
}

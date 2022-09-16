package services

import (
	"ApiOrders/models"
	"ApiOrders/repositories"

	"github.com/gin-gonic/gin"
)

type ServiceImpl struct {
	rr repositories.OrderApi
}

type OrderServiceApi interface {
	GetOrderService(c *gin.Context) gin.H
	CreateOrderService(c *gin.Context) gin.H
	DeleteOrderService(c *gin.Context) gin.H
}

func (s ServiceImpl) GetOrderService(c *gin.Context) gin.H {
	res, err := s.rr.GetOrder(c)
	if err != nil {
		c.JSON(500, "internal server error")
	}
	result := gin.H{
		"result": res,
	}
	return result
}

func (s ServiceImpl) CreateOrderService(c *gin.Context) gin.H {

	requestBody := models.OrderRequestBody{}
	requestItemBody := []models.ItemRequestBody{}

	c.BindJSON(&requestBody)
	c.BindJSON(&requestItemBody)

	res, err := s.rr.CreateOrder(c, requestBody, requestItemBody)
	if err != nil {
		c.JSON(500, "internal server error")
	}
	result := gin.H{
		"result": res,
	}
	return result
}


func (s ServiceImpl) DeleteOrderService(c *gin.Context) gin.H {

	order_id := c.Param("order_id")

	err := s.rr.DeleteOrder(c, order_id)
	if err != nil {
		c.JSON(500, "internal server error")
	}
	result := gin.H{
		"result": "successfully deleted data",
	}
	return result
}

func ProvideService(rr repositories.OrderApi) *ServiceImpl {
	return &ServiceImpl{rr: rr}
}

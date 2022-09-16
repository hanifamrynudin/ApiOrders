package controllers

import (
	"ApiOrders/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB gorm.DB
}

type personController struct {
	ps services.OrderServiceApi
}

func (pc *personController) GetPersonController(c *gin.Context) {
	res := pc.ps.GetOrderService(c)
	c.JSON(200, res)
}

func (pc *personController) CreateOrderController(c *gin.Context) {
	res := pc.ps.CreateOrderService(c)
	c.JSON(200, res)
}

func (pc *personController) DeleteOrderController(c *gin.Context) {
	res := pc.ps.DeleteOrderService(c)
	c.JSON(200, res)
}

func ProvidePersonController(ps services.OrderServiceApi) *personController {
	return &personController{ps: ps}
}

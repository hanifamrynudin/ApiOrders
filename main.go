package main

import (
	"ApiOrders/config"
	"ApiOrders/controllers"
	db "ApiOrders/database"
	"ApiOrders/repositories"
	"ApiOrders/services"

	"github.com/gin-gonic/gin"
)

func main()  {
	
	cfg := config.LoadConfig()
	db := db.InitializeDatabase(cfg.Database.Username, cfg.Database.Password, cfg.Database.Port, cfg.Database.Name, cfg.Database.Host)
	orderRepository := repositories.ProvideRepository(*db)
	orderService := services.ProvideService(orderRepository)
	orderController := controllers.ProvidePersonController(orderService)

	router := gin.Default()

	router.GET("/orders", orderController.GetPersonController)
	router.POST("/orders", orderController.CreateOrderController)
	router.PUT("/orders", orderController.UpdateOrderController)
	router.DELETE("/orders/:order_id", orderController.DeleteOrderController)

	router.Run(cfg.ServerPort)
}
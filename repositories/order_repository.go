package repositories

import (
	"ApiOrders/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB gorm.DB
}

type OrderApi interface {
	GetOrder(c *gin.Context) ([]models.Orders, error)
	CreateOrder(c *gin.Context, requestBody models.OrderRequestBody, requestItemBody []models.ItemRequestBody) (models.Orders, error)
	DeleteOrder(c *gin.Context, id string) (error)
}

func (idb *InDB) GetOrder(c *gin.Context) ([]models.Orders, error) {
	var (
		orders []models.Orders
	)
	err := idb.DB.Model(&models.Orders{}).Preload("Items").Find(&orders).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
	}
	return orders, err
}

func (idb *InDB) CreateOrder(c *gin.Context, requestBody models.OrderRequestBody, requestItemBody []models.ItemRequestBody) (models.Orders, error) {
	var (
		order models.Orders
		item  models.Items
	)

	order.Customer_Name = requestBody.Customer_Name
	order.Ordered_At = requestBody.Ordered_At

	for i := range requestItemBody {
		item.Item_Code = requestItemBody[i].Item_Code
		item.Description = requestItemBody[i].Description
		item.Quantity = requestItemBody[i].Quantity
	}

	err := idb.DB.Create(&models.Orders{
		Customer_Name: order.Customer_Name, 
		Ordered_At: order.Ordered_At, 
		Items: []models.Items{
			{
				Item_Code: item.Item_Code, 
				Description: item.Description, 
				Quantity: item.Quantity}}}).Error

	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
	}
	return order, err
}



func (idb *InDB) DeleteOrder(c *gin.Context, order_id string) (error) {

	var (
		order    models.Orders
	)

	err := idb.DB.First(&order, order_id).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
	}
	
	err = idb.DB.Delete(&order).Error
	if err != nil {
		log.Println("ERROR -> Invalid SQL Syntax")
	}

	return err
}

func ProvideRepository(DB gorm.DB) *InDB {
	return &InDB{DB: DB}
}

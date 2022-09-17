package models

import "time"


type OrderRequestBody struct {
	Order_Id	  uint				`json:"orderId"`
	Customer_Name string			`json:"customerName"`
	Ordered_At    time.Time			`json:"orderedAt"`
	Items         []ItemRequestBody `json:"items"`
}
package models

import "time"


type OrderRequestBody struct {
	Customer_Name string			`json:"customerName"`
	Ordered_At    time.Time			`json:"orderedAt"`
	Items         []ItemRequestBody `json:"items"`
}
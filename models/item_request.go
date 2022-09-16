package models


type ItemRequestBody struct {
	Item_Code   uint		`json:"itemCode"`
	Description string		`json:"description"`
	Quantity    uint		`json:"quantity"`
}
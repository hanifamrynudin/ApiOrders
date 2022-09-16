package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Order_Id      uint `gorm:"primaryKey"`
	Customer_Name string
	Ordered_At    time.Time
	Items         []Items `gorm:"foreignKey:Order_Id"`
}

type Items struct {
	gorm.Model
	Item_Id     uint `gorm:"primaryKey"`
	Item_Code   uint
	Description string
	Quantity    uint
	Order_Id    uint
}

type Person struct {
	Firstname string
	Lastname  string
	Username  string
	Phone     string
	Email     string
	UUID      string
}
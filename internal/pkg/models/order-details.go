package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID      uint
	ProductID    uint
	ProductCount int
	Order        Order
	Product      Product
}

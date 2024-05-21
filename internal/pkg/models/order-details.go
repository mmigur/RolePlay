package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Order     Order
	Product   Product
}

package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId       uint
	TotalPrice   float32
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}

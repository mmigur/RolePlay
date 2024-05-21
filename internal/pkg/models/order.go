package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId     uint
	Products   []Product `gorm:"many2many:order_products;"`
	TotalPrice float32
}

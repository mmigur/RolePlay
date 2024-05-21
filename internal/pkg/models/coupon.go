package models

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	Name        string
	Description string
	Percentage  float32
}

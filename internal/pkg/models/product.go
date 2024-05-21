package models

import "gorm.io/gorm"

type ProductCategory string

const (
	MeatCategory       ProductCategory = "meat"
	VegetablesCategory ProductCategory = "vegetables"
	MilkCategory       ProductCategory = "milk"
	FishCategory       ProductCategory = "fish"
	GroatsCategory     ProductCategory = "groats"
	NutsCategory       ProductCategory = "nuts"
)

type Product struct {
	gorm.Model
	Name              string
	Description       string
	Category          ProductCategory `gorm:"type:product_category_type"`
	IsInStock         bool
	Price             float32
	Weight            float32
	ShelfLife         string
	StorageConditions string
	Brand             string
	Country           string
}

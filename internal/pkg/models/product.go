package models

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
	Id          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	Category    ProductCategory `gorm:"type:product_category_type"`
	Count       int
	Price       float32
	CreatedAt   string
}

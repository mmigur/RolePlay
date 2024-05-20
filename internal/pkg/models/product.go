package models

type ProductType string

const (
	Meat ProductType = "meat"
)

type Product struct {
	Id          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	Category    string
	Count       int
	Price       uint32
	CreatedAt   string
}

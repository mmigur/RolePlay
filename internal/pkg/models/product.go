package models

type Product struct {
	Id                uint `gorm:"primaryKey"`
	ImageUrl          string
	Name              string
	Description       string
	CategoryId        uint `gorm:"not null"`
	IsInStock         bool `gorm:"default:true"`
	OldPrice          float32
	Price             float32
	Weight            float32
	ShelfLife         string
	StorageConditions string
	Brand             string
	Country           string
	CreatedAt         string
}

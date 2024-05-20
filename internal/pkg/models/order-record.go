package models

type OrderRecord struct {
	Id        uint32 `gorm:"primaryKey"`
	UserId    uint32
	Products  []Product
	TotalCost float32
	Date      string
}

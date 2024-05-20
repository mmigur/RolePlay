package models

type User struct {
	Id           uint32 `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	FirstName    string
	MiddleName   string
	LastName     string
	Address      string
	Orders       []OrderRecord
	Coupons      []Coupon
	RegisteredAt string
}

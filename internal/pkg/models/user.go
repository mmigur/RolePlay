package models

type User struct {
	Id           uint32 `gorm:"primaryKey"`
	FirstName    string
	MiddleName   string
	LastName     string
	Nickname     string
	Address      string
	Coupons      []Coupon
	Email        string
	Password     string
	RegisteredAt string
}

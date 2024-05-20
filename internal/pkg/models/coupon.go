package models

type DiscountType int

const (
	OneThousand    DiscountType = 3
	ThreeThousands DiscountType = 10
	Registration   DiscountType = 5
)

type Coupon struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	Description  string
	DiscountType DiscountType `gorm:"type:discount_type"`
	Cost         int
	CreatedAt    string
}

package models

type DiscountType float32

const (
	OneThousand    DiscountType = 0.03
	ThreeThousands DiscountType = 0.1
	Registration   DiscountType = 0.05
)

type Coupon struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	Description  string
	DiscountType DiscountType
	CreatedAt    string
}

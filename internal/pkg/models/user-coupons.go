package models

import "gorm.io/gorm"

type UserCoupon struct {
	gorm.Model
	UserID   uint
	CouponID uint
	User     User
	Coupon   Coupon
}

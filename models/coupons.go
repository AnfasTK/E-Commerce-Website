package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	CouponCode         string `gorm:"unique"`
	FixedDiscount      int
	DiscountPercentage float64
	MaxDiscount        int
	MinProductPrice    float64
	UsersUsedCount     int
	MaxUseCount        int
	ValidFrom          time.Time
	ExpirationDate     time.Time
	IsFixedCoupon      bool   `gorm:"default:false"`
	IsActive           bool   `gorm:"default:true"`
	CouponType         string `gorm:"default:'percentage'"`
	Status             string
}

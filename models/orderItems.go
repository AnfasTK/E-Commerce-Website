package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID          uint
	UserID           uint
	ProductID        uint
	Quantity         int
	Subtotal         float64
	OrderStatus      string `gorm:"size:50"`
	IsDelivered      bool   `gorm:"default:false"`
	DeliveryDate     time.Time
	ReturnableStatus bool `gorm:"default:true"`
	ReturnDate       time.Time
	OrderDetail      Order
	UserAuth         UserAuth
	ProductDetail    ProductDetail
}
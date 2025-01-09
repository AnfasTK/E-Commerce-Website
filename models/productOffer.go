package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductOffer struct {
	gorm.Model
	OfferName       string  `gorm:"size:255" json:"offername"`
	OfferDetails    string  `gorm:"size:255" json:"offer"`
	OfferAmount     int     `json:"discount"`
	OfferPercentage float64 `json:"discountpercentage"`
	StartDate       time.Time
	EndDate         time.Time
	ProductID       uint `json:"productid"`
	Product         ProductDetail
	IsValid         bool `gorm:"default:false"`
}

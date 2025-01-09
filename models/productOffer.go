package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductOffer struct {
	gorm.Model
	OfferName       string `gorm:"size:255"`
	OfferDetails    string `gorm:"size:255"`
	OfferAmount     int
	OfferPercentage float64
	StartDate       time.Time
	EndDate         time.Time
	ProductID       uint
	IsValid         bool `gorm:"default:false"`
}

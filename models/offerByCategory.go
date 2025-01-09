package models

import (
	"time"

	"gorm.io/gorm"
)

type OfferByCategory struct {
	gorm.Model
	CategoryOfferName       string `gorm:"size:255"`
	CategoryOfferPercentage float64
	OfferDescription        string `gorm:"size:255"`
	CategoryID              uint
	OfferStatus             bool `gorm:"default:false"`
	StartDate               time.Time
	EndDate                 time.Time
	IsOfferDeleted          bool `gorm:"default:false"`
	Category                Category
}

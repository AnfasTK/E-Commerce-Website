package models

import (
	"gorm.io/gorm"
)

type ProductDescription struct {
	gorm.Model
	ProductID          uint
	DescriptionHeading string        `gorm:"size:255" json:"descriptionheading"`
	Description        string        `gorm:"size:1000" json:"description"`
	ProductDetail      ProductDetail `gorm:"foreignKey:ProductID"`
}

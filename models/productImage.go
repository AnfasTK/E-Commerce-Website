package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductImages  string        `gorm:"not null"`
	ProductID      uint          `gorm:"not null"`
	Product        ProductDetail `gorm:"foreignKey:ProductID"`
}

package models

import "gorm.io/gorm"

type ProductVariantsImage struct {
	gorm.Model
	ProductVariantsImages string                `gorm:"not null"`
	ProductVariantID      uint                  `gorm:"not null"`
	ProductVariant        ProductVariantDetails `gorm:"foreignKey:ProductVariantID"`
}

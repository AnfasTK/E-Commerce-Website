package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductImages  []byte        `gorm:"type:bytea"`
	ProductBanners []byte        `gorm:"type:bytea"`
	ProductPosters []byte        `gorm:"type:bytea"`
	ProductID      uint          `gorm:"not null"`
	Product        ProductDetail `gorm:"foreignKey:ProductID"`
}

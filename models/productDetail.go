package models

import (
	"gorm.io/gorm"
)

type ProductDetail struct {
	gorm.Model
	ProductName  string               `gorm:"size:255" json:"productname"`
	CategoryID   uint
	BrandName    string               `gorm:"size:100" json:"brand"`
	IsReturnable bool                 `gorm:"default:true"`
	Category     Categories           `gorm:"foreignKey:CategoryID"`
	Descriptions []ProductDescription `gorm:"foreignKey:ProductID"`
	Variants     []ProductVariant     `gorm:"foreignKey:ProductID"`
	Images       []ProductImage       `gorm:"foreignKey:ProductID"`
	Offers       []ProductOffer       `gorm:"foreignKey:ProductID"`
}

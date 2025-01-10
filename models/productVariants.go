package models

import (
	"gorm.io/gorm"
)

type ProductVariant struct {
	gorm.Model
	ProductID      uint                   `gorm:"not null"`
	Size           string                 `gorm:"size:100" json:"size"`
	Colour         string                 `gorm:"size:100" json:"color"`
	Ram            string                 `gorm:"size:100" json:"ram"`
	Storage        string                 `gorm:"size:100" json:"storage"`
	StockQuantity  int                    `json:"stock"`
	RegularPrice   float64                `json:"regular"`
	SalePrice      float64                `json:"saleprice"`
	SKU            string                 `gorm:"unique" json:"sku"`
	ProductSummary string                 `gorm:"size:255" json:"summery"`
	IsDeleted      bool                   `gorm:"default:false"`
	Product        ProductDetail          `gorm:"foreignKey:ProductID"`
	Specification  []ProductSpecification `gorm:"foreignKey:ProductVariantsID"`
}

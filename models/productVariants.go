package models

import (
	"gorm.io/gorm"
)

type ProductVariant struct {
	gorm.Model
	ProductID      uint
	Size           string `gorm:"size:100"`
	Color          string `gorm:"size:100"`
	RAM            string `gorm:"size:100"`
	Storage        string `gorm:"size:100"`
	StockQuantity  int
	RegularPrice   float64
	SalePrice      float64
	SKU            string `gorm:"unique"`
	ProductSummary string `gorm:"size:255"`
	IsDeleted      bool   `gorm:"default:false"`
	ProductDetail  ProductDetail
}

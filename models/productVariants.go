package models

import (
	"gorm.io/gorm"
)

type ProductVariant struct {
	gorm.Model
	ProductID      uint
	Size           string  `gorm:"size:100" json:"size"`
	Color          string  `gorm:"size:100" json:"color"`
	RAM            string  `gorm:"size:100" json:"ram"`
	Storage        string  `gorm:"size:100" json:"storage"`
	StockQuantity  int     `json:"stock"`
	RegularPrice   float64 `json:"regular"`
	SalePrice      float64 `json:"salefrice"`
	SKU            string  `gorm:"unique" json:"sku"`
	ProductSummary string  `gorm:"size:255" json:"summery"`
	IsDeleted      bool    `gorm:"default:false"`
	ProductDetail  ProductDetail
}

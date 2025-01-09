package models

import (
	"gorm.io/gorm"
)

type ProductDetail struct {
	gorm.Model
	ProductName  string `gorm:"size:255"`
	CategoryID   uint
	BrandName    string `gorm:"size:100"`
	IsReturnable bool   `gorm:"default:true"`
	Category     Category
}

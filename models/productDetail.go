package models

import (
	"gorm.io/gorm"
)

type ProductDetail struct {
	gorm.Model
	ProductName  string `gorm:"size:255" json:"productname"`
	CategoryID   uint
	BrandName    string `gorm:"size:100" json:"brand"`
	IsReturnable bool   `gorm:"default:true"`
	Category     Categories
}

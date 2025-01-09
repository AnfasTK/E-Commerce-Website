package models

import "gorm.io/gorm"

type ProductSpecification struct {
	gorm.Model
	ProductSummary     string `gorm:"size:255"`
	ProductVariantsID  uint
	SpecificationKey   string `gorm:"size:255"`
	SpecificationValue string `gorm:"size:255"`
	ProductVariant     ProductVariant
}

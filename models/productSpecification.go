package models

import "gorm.io/gorm"

type ProductSpecification struct {
	gorm.Model
	ProductSummary     string         `gorm:"size:255" json:"summery"`
	ProductVariantsID  uint           `gorm:"not null"`
	SpecificationKey   string         `gorm:"not null;size:255" json:"specificationkey"`
	SpecificationValue string         `gorm:"not null;size:255" json:"specificationvalue"`
	ProductVariant     ProductVariant `gorm:"foreignKey:ProductVariantsID"`
}

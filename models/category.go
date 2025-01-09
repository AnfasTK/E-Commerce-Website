package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"size:255"`
	Status bool   `gorm:"default:false"`
}

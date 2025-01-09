package models

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	UserID        uint
	ProductID     uint
	Value         float64
	UserAuth      UserAuth
	ProductDetail ProductDetail
}

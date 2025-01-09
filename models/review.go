package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID        uint
	ProductID     uint
	Review        string `gorm:"size:1000"`
	UserAuth      UserAuth
	ProductDetail ProductDetail
}

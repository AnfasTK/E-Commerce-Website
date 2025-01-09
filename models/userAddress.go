package models

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	FirstName string `gorm:"size:100"`
	LastName  string `gorm:"size:100"`
	Mobile    string `gorm:"size:15"`
	Address   string `gorm:"type:text"`
	Landmark  string `gorm:"size:255"`
	Country   string `gorm:"size:100"`
	State     string `gorm:"size:100"`
	City      string `gorm:"size:100"`
	PinCode   int    `gorm:"size:6"`
	UserID    uint
	IsDefault bool `gorm:"default:false"`
	UserAuth  UserAuth
}

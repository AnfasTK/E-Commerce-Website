package models

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	FirstName string   `gorm:"size:100" json:"user_firstname"`
	LastName  string   `gorm:"size:100" json:"user_lastname"`
	Mobile    string   `gorm:"size:15" json:"user_number"`
	Address   string   `gorm:"type:text" json:"user_address"`
	Landmark  string   `gorm:"size:255" json:"user_landmark"`
	Country   string   `gorm:"size:100" json:"user_country"`
	State     string   `gorm:"size:100" json:"user_state"`
	City      string   `gorm:"size:100" json:"user_city"`
	PinCode   string   `gorm:"not null" json:"user_pincode"`
	UserID    uint     `gorm:"not null"`
	IsDefault bool     `gorm:"default:false"`
	UserAuth  UserAuth `gorm:"foreignKey:UserID"`
}

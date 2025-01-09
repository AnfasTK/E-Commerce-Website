package models

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model
	UserID     uint   `gorm:"unique"`
	Mobile     string `gorm:"unique;size:15"`
	Country    string `gorm:"size:100"`
	State      string `gorm:"size:100"`
	Pincode    string `gorm:"size:10"`
	ProfilePic []byte `gorm:"type:bytea"`
	UserAuth   UserAuth
}

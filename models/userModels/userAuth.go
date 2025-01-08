package models

import (
	"time"

	"gorm.io/gorm"
)

type UserAuth struct {
	gorm.Model
	FullName    string `gorm:"not null"`
	Email       string `gorm:"unique,not null"`
	Passowrd    string `gorm:"not null"`
	IsBlocked   bool   `gorm:"default:false"`
	IsDeleted   bool   `gorm:"default:false"`
	LastLoginAt time.Time
}

package models

import (
	"gorm.io/gorm"
	"time"
)

type Otp struct {
	gorm.Model
	Email       string `gorm:"size:255"`
	OTP         string `gorm:"size:10"`
	CreatedTime time.Time
	ExpireTime  time.Time
	UserAuth    UserAuth
}

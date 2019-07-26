package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	WxOpenID      string    `json:"WxOpenID"`
	WxSessionKey  string    `json:"WxSessionKey"`
	WxLocalToken  string    `json:"WxLocalToken"`
	LastLoginDate time.Time `json:"LastLoginDate"`
	Name          string    `json:"Name"`
	SchoolID      string    `json:"SchoolID"`
}

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//ReadLog 打卡记录
type ReadLog struct {
	gorm.Model
	UserID uint      `json:"UserID"`
	Time   time.Time `json:"Time"`
}

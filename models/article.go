package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Article 文章
type Article struct {
	gorm.Model
	Time    time.Time `json:"Time"`
	Content string    `json:"Content" gorm:"size:10000"`
}

package models

import (
	"server/config"

	"github.com/jinzhu/gorm"

	//mysql dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", config.DbUser+":"+config.DbPassword+"@/"+config.DbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("fail to open database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&ReadLog{})
}

// Db 获取db
func Db() *gorm.DB {
	return db
}

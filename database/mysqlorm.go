package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlOrm *gorm.DB

func init() {

	dns := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	MysqlOrm = db
}

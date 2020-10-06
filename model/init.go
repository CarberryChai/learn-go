package model

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局唯一数据库链接
var DB *gorm.DB

// Init 初始化数据库
func Init() {
	mysqlConfig := os.Getenv("MYSQL_CONF")
	db, err := gorm.Open(mysql.Open(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&User{})
	DB = db
}

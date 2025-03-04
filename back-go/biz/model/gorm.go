package model

import (
	"fmt"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func GetDB() *gorm.DB {

	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.Debug()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	runtime.KeepAlive(DB)
	return DB
}

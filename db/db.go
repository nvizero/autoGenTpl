package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user = "root"
	pwd  = "root"
	host = "127.0.0.1"
	port = 3306
	db   = "livewire"
)

// 创建数据库，如果存在则删除并重新创建
func CreateOrRecreateDB(dbName string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, pwd, host, port)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 创建数据库，如果存在的话
	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " CHARACTER SET utf8 COLLATE utf8_general_ci").Error; err != nil {
		panic(err)
	}
}

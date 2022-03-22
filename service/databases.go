package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//var DbService *gorm.DB
//
//func init() {
//	DbService = mysqlConnection()
//}

func MysqlConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/dj?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic("数据库连接失败")
	}
	db.DB().SetConnMaxLifetime(60 * time.Second)
	return db
}

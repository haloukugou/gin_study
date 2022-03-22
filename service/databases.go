package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"os"
	"time"
)

//var DbService *gorm.DB
//
//func init() {
//	DbService = mysqlConnection()
//}

func MysqlConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/dj1?charset=utf8mb4&parseTime=true")
	if err != nil {
		gin.DisableConsoleColor()
		nowTime := time.Now().Format("2006-01-02")
		f, _ := os.Create("../log/" + nowTime + ".log")
		gin.DefaultWriter = io.MultiWriter(f)
		panic("数据库连接失败")
	}
	db.DB().SetConnMaxLifetime(60 * time.Second)
	return db
}

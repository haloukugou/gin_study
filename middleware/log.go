package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		gin.DisableConsoleColor()
		nowTime := time.Now().Format("2006-01-02")
		logFile, err := os.Create("../log/" + nowTime + ".log")
		if err != nil {
			fmt.Println("could not create log file")
		}
		//fmt.Println(reflect.TypeOf(f))
		gin.DefaultWriter = io.MultiWriter(logFile)
		//status := c.Writer.Status()
		//if 200 != status {
		//
		//}
		c.Next()
	}
}

package main

import (
	"gin/controller"
	"gin/middleware"
	"github.com/gin-gonic/gin"
)

// gin的helloWorld

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	r.Use(middleware.Log())
	// 路由组1 ，处理GET请求
	v1 := r.Group("/v1")
	{
		v1.GET("/echo", controller.Login)
		v1.POST("/login", controller.UserLogin)
	}

	r.Run(":8000")
}

package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS bool = true
	FAILED  bool = false
)

func ResponseRes(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, gin.H{"code": code, "msg": msg, "data": data})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"state": SUCCESS,
		"msg":   "success",
		"data":  data,
	})
}

package controller

import (
	"fmt"
	"gin/common"
	"gin/model"
	"gin/service"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {

}

func Login(c *gin.Context) {
	//name := c.DefaultQuery("name", "jack")
	//c.String(200, fmt.Sprintf("hello %s\n", name))
	// todo 年月日时分秒
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	// todo 年月日
	//nowTime := time.Now().Format("2006-01-02")

	c.String(200, fmt.Sprintf(nowTime))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func UserLogin(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("pwd")
	//c.String(http.StatusOK, fmt.Sprintf("account:%s,pwd:%s", account, pwd))
	//return
	if "" == account {
		common.ResponseRes(c, 4001, "账号为空", []int{})
		return
	}
	if "" == pwd {
		common.ResponseRes(c, 4001, "密码为空", []int{})
		return
	}
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(nowTime)
	//res := struct {
	//	Account  string `json:"account"`
	//	Password string `json:"password"`
	//}{
	//	Account:  account,
	//	Password: pwd,
	//}
	//common.ResponseRes(c, 2000, "success1", res)
	var user model.User
	db := service.MysqlConnection()
	db.First(&user, 1)
	common.ResponseRes(c, 2000, "success", user)
	//fmt.Println(user)
}

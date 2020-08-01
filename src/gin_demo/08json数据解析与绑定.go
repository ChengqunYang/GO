package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	客户端传参,后端接收并解析到结构体
*/
// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段,若接收值为空,则报错,是必须的字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 1.创建路由
	engine := gin.Default()
	// 2.json绑定
	engine.POST("loginJSON", func(context *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将request的body中的数据,自动按照json格式解析到结构体
		if err := context.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H 封装了生成json数据的工具
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "root" || json.Password != "admin" {
			context.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	engine.Run(":8080")
}

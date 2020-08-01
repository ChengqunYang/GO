package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 表单传输为post请求,http常见的传输格式为四种,表单参数
//可以通过PostForm()方法获取,该方法默认解析的是application/x-www-form-urlencoded或者multipart/form-data格式的参数
func main() {
	// 1.创建路由
	engine := gin.Default()
	engine.POST("/form", func(context *gin.Context) {
		// 表单参数设置默认值
		typel := context.DefaultPostForm("type", "alert")
		// 接受其他的
		username := context.PostForm("username")
		password := context.PostForm("password")
		// 多选框
		hobbys := context.PostFormArray("hobby")
		context.String(http.StatusOK, fmt.Sprintf("type is %s username is %s,password is %s,hobby is %v", typel, username, password, hobbys))
	})
	engine.Run()
}

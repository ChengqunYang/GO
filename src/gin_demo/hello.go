package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 的hello world

func main() {
	// 1.创建路由
	// 默认使用了两个中间件Logger(), Recovery()
	engine := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	engine.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	engine.PUT("/xxxPut", geting)
	engine.POST("/xxxPost")
	// 3.监听端口，默认8080
	engine.Run()
}

// 函数必须有一个*gin.Context参数
func geting(c *gin.Context) {

}

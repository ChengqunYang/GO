package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// routes group 是为了管理一些相同的url
func main() {
	// 1.创建路由
	engine := gin.Default()
	// 路由组:处理url请求
	v1 := engine.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := engine.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("submit", submit)
	}
	engine.Run(":8080")
}
func login(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "jack")
	ctx.String(200, fmt.Sprintf("hello %s\n", name))
}
func submit(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "lily")
	ctx.String(200, fmt.Sprintf("hello %s\n", name))
}

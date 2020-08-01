package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/currency"
)

// gin框架中采用的路由库是基于httprouter做的
/*
	httprouter会将所有的路由规则构造一颗前缀树
	例如有root and as  at cn com
*/

/*
				/
	s					blog	/contact	/about
earch	support		:post

*/
func main() {
	// 1.创建路由
	engine := gin.Default()
	engine.POST("/", xxx)
	engine.POST("search", xxx)
	engine.POST("support", xxx)
	engine.POST("/blog/:post", xxx)
	engine.POST("/contact", xxx)
	engine.POST("/about", xxx)
	engine.Run(":8080")
}

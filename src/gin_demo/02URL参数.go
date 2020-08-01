package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 可以使用DefaultQuery()或者Query()方法获取URL参数
// DefaultQuery()若参数不对，则返回默认值，Query()若不存在，返回空串
//API ? name = zhangsan
func main() {
	engine := gin.Default()
	engine.GET("/welcome", func(context *gin.Context) {
		// DefaultQuery()第二个参数是默认值
		name := context.DefaultQuery("name", "Jack")
		context.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})
	engine.Run()
}

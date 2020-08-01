package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 可以通过Context的Param方法来获取API参数
func main() {
	// 1.创建路由
	engine := gin.Default()
	// api参数
	//engine.GET("/user/:name", func(context *gin.Context) {
	//	name := context.Param("name")
	//	context.String(http.StatusOK,name)
	//})

	engine.GET("/user/:name/*action", func(context *gin.Context) {
		//
		name := context.Param("name")
		action := context.Param("action")
		context.String(http.StatusOK, name+" is "+action)
	})
	engine.Run()

}

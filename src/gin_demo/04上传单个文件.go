package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// application/form-data格式用于文件上传
// gin文件上传与原生的net/http方法类似，不同之处在于gin把原生的request封装到context.Request中

func main() {
	// 1.创建路由
	engine := gin.Default()
	engine.POST("/upload", func(context *gin.Context) {
		// 表单取文件
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		// 默认保存到项目的根目录，名字就用本身的
		_ = context.SaveUploadedFile(file, file.Filename)
		// 打印信息
		context.String(200, fmt.Sprintf("%s upload!", file.Filename))

	})
	engine.Run()
}

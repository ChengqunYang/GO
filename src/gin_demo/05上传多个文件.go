package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	engine := gin.Default()
	// 限制上传大小 8MB,默认值为32MB
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/uploads", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"] //files为html里元素的name
		// 遍历所有图片
		for _, file := range files {
			// 逐个存储
			if err := context.SaveUploadedFile(file, file.Filename); err != nil {
				context.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		context.String(200, fmt.Sprintf("update ok %d files", len(files)))
	})
	engine.Run()
}

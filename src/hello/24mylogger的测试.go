package main

import (
	"mylogger"
	"time"
)
//声明一个全局的接口变量
var log mylogger.Logger
//测试我们自己写的日志库
func main() {
	log = mylogger.NewConsoleLogger("Debug") //终端日志实例
	log = mylogger.NewFileLogger("waring","./","test.log",10*1024) //日志文件实例
	for true {
		id := 10011
		name := "张三"
		log.Debug("这是一条Debug日志,id: %d  name: %s", id, name)
		log.Info("这是一条info日志")
		log.Waring("这是一条waring日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 2)
	}
}

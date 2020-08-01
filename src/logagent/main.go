package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logagent/kafka"
	"logagent/taillog"
	"time"
)
// 0. 加载配置文件
var(
	cfg, err = ini.Load("./conf/config.ini")
)

func run() {
	// 1. 读取日志
	for {
		select {
			case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SendToKafka(cfg.Section("kafka").Key("topic").String(),line.Text)
		default:
			time.Sleep(time.Second)
}
	}

}
//logagetn入口程序
func main() {
	if err!=nil {
		fmt.Printf("load ini failed,err: %v\n",err)
	}
	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.Section("kafka").Key("address").String()})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n",err)
		return
	}
	fmt.Println("init kafka success.")
	// 2. 打开日志文件准备日志
	err = taillog.Init(	cfg.Section("taillog").Key("path").String())
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n",err)
		return
	}
	fmt.Println("init taillog success.")
	run()
}

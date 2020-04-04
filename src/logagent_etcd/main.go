package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logagent_etcd/etcd"
	"logagent_etcd/kafka"
	"logagent_etcd/taillog"
	"time"
)

// 0. 加载配置文件
var (
	cfg, err = ini.Load("./conf/config.ini")
)

func run() {
	// 1. 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SendToKafka(cfg.Section("kafka").Key("topic").String(), line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}

//logagetn入口程序
func main() {
	if err != nil {
		fmt.Printf("load ini failed,err: %v\n", err)
	}
	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.Section("kafka").Key("address").String()})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")
	//2. 初始化etcd
	duration, err1 := cfg.Section("etcd").Key("timeout").Duration()
	if err1 != nil {
		fmt.Printf("get duration failed,err:%v\n", err1)
		return
	}
	err = etcd.Init(cfg.Section("etcd").Key("address").String(), duration)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")
	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.Section("etcd").Key("collect_log_key").String())
	// 2.2 派一个哨兵去监视日志收集项的变化(有变化及时通知logagent实现热配置)
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success,%v\n", logEntryConf)
	// 2. 打开日志文件准备日志
	/*err = taillog.Init(	cfg.Section("taillog").Key("path").String())
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n",err)
		return
	}
	fmt.Println("init taillog success.")
	run()*/

}

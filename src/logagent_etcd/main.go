package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logagent_etcd/etcd"
	"logagent_etcd/kafka"
	"logagent_etcd/taillog"
	"sync"
)

// 0. 加载配置文件
var (
	cfg, err = ini.Load("./conf/config.ini")
)


//logagetn入口程序
func main() {
	if err != nil {
		fmt.Printf("load ini failed,err: %v\n", err)
	}
	// 1. 初始化kafka连接
	maxSize, err2 := cfg.Section("kafka").Key("chan_max_size").Int()
	if err2 != nil {
		fmt.Printf("get maxsize failed,err:%v\n",err2)
	}
	err = kafka.Init([]string{cfg.Section("kafka").Key("address").String()},maxSize)
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
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success,%v\n", logEntryConf)

	for index, value := range logEntryConf {
		fmt.Printf("index:%v,value:%v\n",index,value)
	}
	// 3.收集日志发往kafka    派一个哨兵去监视日志收集项的变化(有变化及时通知logagent实现热配置)
	taillog.Init(logEntryConf)
	newConfChan := taillog.NewConfChan()  // 从taillog中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(cfg.Section("etcd").Key("collect_log_key").String(),newConfChan) // 哨兵发现最新的配置信息会通知上面的那个通道
	wg.Wait()

}

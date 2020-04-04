package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// etcd watch
func main() {
	// 1.建立连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// watch
	// 派一个小弟去监视这个key的变化(新增,修改,删除),只要有消息就通过通道传话返回一个只读的channel
	watch := cli.Watch(context.Background(), "ycq")
	// 从通道尝试取值(监视的信息)
	for wresp := range watch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}

}

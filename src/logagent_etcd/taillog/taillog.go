package taillog

/*
	专门从文件收集日志模块
*/
import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"logagent_etcd/kafka"
)

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了能实现退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      "topic",
		instance:   nil,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init() // 根据路径去打开对应的日志
	return
}

var err error

func (t *TailTask) init() {
	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
	}
	go t.run() //直接去采集日志发送到kafka,当run函数退出的时候goroutine也就结束了
}

func (t *TailTask) run() {
	for true {
		// 从tailObj的通道中一行一行的读取日志内容
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s 结束了...\n",t.path,t.topic)
			return
		case line := <-t.instance.Lines:
			// 3.2发送到kafka
			//kafka.SendToKafka(t.topic,line.Text) // 函数调用函数,只有当内层函数执行完才可以继续执行外层for

			// 先把日志数据发送到一个通道中
			kafka.SendToChan(t.topic, line.Text)
			// kafka那个包中有单独的goroutine去取日志数据发送到kafaka

		}
	}
}

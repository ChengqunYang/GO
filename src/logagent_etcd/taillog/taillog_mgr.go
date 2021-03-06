package taillog

import (
	"fmt"
	"logagent_etcd/etcd"
	"time"
)

var (
	tskMgr *tailLogMgr
)

// tailTask 管理者
type tailLogMgr struct {
	logEntry    [] *etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), //无缓冲区的通道
	}
	// 遍历出每一个配置项
	for _, logEntry := range logEntryConf {
		// conf:etcd.LogEntry
		// logEntry.Path: 要收集的日志文件的路径
		// 用获取到的每一个日志项建立一个tail任务,初始化的时候起了多少个task得记录下来,为了后续判断方便
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}
	go tskMgr.run()
}

//监听自己的newConfChan,有了新的配置过来之后就做对应的处理
func (t *tailLogMgr) run() {
	for true {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有,不需要操作
					continue
				} else {
					// 如果没有,就是新增的
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			// 找出原来有,现在没有了的,要删掉
			for _, c1 := range t.logEntry {   // 从原来的配置中依次拿出配置项
				isDelete := true
				for _, c2 := range newConf {  // 去新的配置中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}

			// 2. 配置删除
			fmt.Println("新的配置来了!", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 一个函数,向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}

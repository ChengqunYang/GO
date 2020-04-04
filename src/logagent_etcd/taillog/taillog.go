package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

//专门从日志文件收集日志的模块

func Init(fileName string) (err error) {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的那个地方开始读
		ReOpen:    true,                                 //重新打开
		MustExist: false,                                //文件不存在会报错
		Poll:      true,                                 //
		Follow:    true,                                 //是否跟随
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	return
}
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}

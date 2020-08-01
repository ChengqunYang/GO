package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tail的用法示例
func main() {
	fileName := "./my.log"
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的那个地方开始读
		ReOpen:    true,                                 //重新打开
		MustExist: false,                                //文件不存在会报错
		Poll:      true,                                 //
		Follow:    true,                                 //是否跟随
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}

	var (
		line *tail.Line
		ok   bool
	)
	for true {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Print("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}

}

package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("usage:xxx file")
		return
	}

	fileName := list[1]
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("fileName = ", info.Name())
	fmt.Println("size = ", info.Size())
	/*
		结果:
			go run 03获取文件属性.go D:\demo.text
		fileName =  demo.text
		size =  61
	*/
}

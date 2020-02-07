package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//从d盘的demo.txt复制文件内容到项目目录下的demo.txt中
	sf, err := os.Open("d:/demo.text")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	df, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	buf := make([]byte, 4*1024) //创建一个4k大小的缓冲区
	for {
		n, err := sf.Read(buf) //从源文件读取内容,n为读取到的内容的数目
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}
		//往目的文件中写内容,读多少写多少
		df.Write(buf[:n])
	}
	//操作完毕,需要关闭文件
	defer sf.Close()
	defer df.Close()

	//io.Copy(df,sf) //使用该函数可直接复制源文件到目的文件中

}

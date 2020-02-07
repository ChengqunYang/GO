package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//1.设备文件   屏幕,键盘等
//2.磁盘文件    文本文件,二进制文件

//写文件
func WriteFile(path string) {
	//新建文件并打开
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	} else {
		//往文件中追加数据
		for i := 0; i < 10; i++ {
			buf := fmt.Sprintf("i = %d\n", i)
			_, err := f.WriteString(buf)
			if err != nil {
				fmt.Println("err = ", err)
				return
			}
		}
	}
	//使用完毕,需要关闭文件
	defer f.Close()
}

//通过bufio往文件中追加写内容
func WriteFileBufio(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello go go go!")
	}
	writer.Flush()
}

func WriteFileIoUtil(path string) {
	str := "hello world!"
	err := ioutil.WriteFile(path, []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}
}

//读文件
//一次读一个自己指定的字节数组大小
func ReadFile(path string) {
	//以只读的方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//读取文件
	buf := make([]byte, 1024*2) //2k

	//n代表从文件读取内容的长度
	n, err := f.Read(buf)
	if err != nil && err != io.EOF { //文件出错，同时没有到结尾（eof代表结束符）
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))

	//关闭文件
	defer f.Close()
}

//借助bufio实现一次读取一行
func ReadOneLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//新建一个缓冲区,把内容放在缓冲区中
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n') //会将\n也读取进去,即自动换行
		if err != nil {
			if err == io.EOF { //表示文件已经读取到末尾
				break
			}
			fmt.Println("err = ", err)
			return
		}
		fmt.Print(string(buf))
	}
	//关闭文件
	defer f.Close()
}

func ReadFileIoutil(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	fmt.Println(string(content))
}
func main() {
	path := "d:/demo.text"
	//新建文件
	//打开文件

	//读文件
	//ReadFile(path)
	//ReadOneLine(path)
	//ReadFileIoutil(path)

	//写文件
	//WriteFile(path)
	//WriteFileBufio(path)
	WriteFileIoUtil(path)

	//删除文件
}

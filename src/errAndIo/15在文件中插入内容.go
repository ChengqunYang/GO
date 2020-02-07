package main

import (
	"fmt"
	"io"
	"os"
)

func f2() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./demo1.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}

	//因为没有办法直接在一个文件中间插入内容,所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./tmp.txt", os.O_RDWR, 0644) //0644表示要打开的文件的类型
	if err != nil {
		fmt.Println("create tmp file failed,err", err)
		return
	}

	//读取文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println("read from file filed,err:", err)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])
	//再写入要插入的内容
	var s []byte
	s = []byte{'f'}
	tmpFile.Write(s)
	//紧接着将源文件后续的内容写入临时文件
	var x [1024]byte
	for true {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Println("read from failed ,err:", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./tmp.txt", "./demo1.txt")

	/*//seek函数用来移动光标
	fileObj.Seek(3,0) //光标移动  第一个变量表示光标移动几个位置,第二个变量表示以什么为基准
	//fileObj.Write([]byte{'f'})//往光标所在位置插入一个字符f,但是后面的字符并不会向后移位*/

}
func main() {
	f2()
}

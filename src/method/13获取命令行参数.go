package main

//所有用go语言编译的可执行程序都必须有一个名叫main的包,一个可执行程序有且只有一个main包
import (
	"fmt"
	"os"
)

func main() {
	//接受用户传递的参数,都是以字符串格式传递的
	list := os.Args
	n := len(list)
	fmt.Println("n=", n)

	//打印这些参数
	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}

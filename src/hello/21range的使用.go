package main

import "fmt"

func main() {
	str := "abc"
	//通过for循环打印每一个字符
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		fmt.Printf("%c\n", str[i])
	}

	//使用迭代打印每个元素,默认返回两个值
	for i, data := range str {
		fmt.Printf("str[%d] : %c\n", i, data)
	}

	for i := range str {
		fmt.Printf("%d\n", i)
		fmt.Printf("%c\n", str[i])
	}
}

package main

import "fmt"

func main() {
	var a int
	fmt.Println("请输入变量a:")
	//阻塞等待用户输入
	fmt.Scanf("%d", &a)
	fmt.Printf("a = %d\n", a)

	var b int
	fmt.Println("请输入变量b")
	//自动匹配格式
	fmt.Scan(&b)
	fmt.Printf("b = %d", b)
}

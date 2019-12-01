package main

import "fmt"

func main() {
	a := "王思聪"

	//if后面就是条件
	if a == "王思聪" { //左括号不能换行
		fmt.Println("这个人就是王思聪")
	}

	//if支持一个初始化语句,初始化语句和判断条件以分号分隔
	if a := 10; a == 10 {
		fmt.Println("a的值就是", a)
	}
	if true && false {
		fmt.Println("出错了")
	} else {
		fmt.Println("正确的")
	}

	b := 19
	if b > 19 {
		fmt.Println("a是大于十九的")
	} else if b < 19 {
		fmt.Println("a是小于十九的")
	} else if b == 19 {
		fmt.Println("a是等于十九的")
	}
}

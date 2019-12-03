package main

import "fmt"

func test() {
	a := 10
	fmt.Printf("a=%d", a)
}
func main() {
	var d int
	//定义在{}里的变量就是局部变量,只能在{}里面有效
	//作用域,变量其作用的范围

	//错误,a只在test那个函数中有效
	//a = 111
	{
		i := 10
		fmt.Println("i = ", i)
	}
	//错误,i只在它所在的那个{}内有效
	//i =111;
	if flag := 3; flag >= 3 {
		d = 12
		fmt.Println(flag)
		fmt.Println(d)
	}
	//错误,flag只在if那个{}内有效
	//flag = 19
}

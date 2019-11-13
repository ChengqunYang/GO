package main

import "fmt"

//go语言可以返回多个值,此函数返回a,b,c
func test() (a, b, c int) {
	return 1, 2, 3
}
func main() {
	/*	a:= 10
		b:= 20
		c:= 30*/
	a, b := 10, 20
	fmt.Println(a, b)
	a, b = b, a //交换两个变量的值
	fmt.Println(a, b)

	//_:匿名变量  丢弃数据,不处理
	tmp, _ := 7, 8
	fmt.Println(tmp)

	//如果只想获取返回值的后两个,则可以通过匿名变量的方式来获取第一个变量
	_, c, d := test()
	fmt.Println(c, d)
}

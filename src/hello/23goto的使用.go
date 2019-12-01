package main

import "fmt"

//goto无条件跳转
func main() {
	//break只能用在loop,switch,select
	//continue 只能用在loop中
	//goto可以用在任何地方,但是不能跨函数使用
	goto End
	fmt.Println("222222222")
End:
	fmt.Println("33333333")
}

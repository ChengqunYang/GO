package main

import "fmt"

//值传递
func swap(a, b int) {
	a, b = b, a
	fmt.Printf("a= %d,b = %d\n", a, b)
}
func main() {
	var a int = 10
	var b int = 20
	//通过一个函数,交换a和b的值
	swap(a, b)
	fmt.Printf("a= %d,b = %d\n", a, b)
}

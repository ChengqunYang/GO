package main

import "fmt"

//指针作为变量,地址传递
func swap1(a, b *int) {
	*a, *b = *b, *a
	fmt.Printf("a= %d,b = %d\n", *a, *b)

}
func main() {
	a, b := 10, 20
	//通过一个函数来交换a和b的值
	swap1(&a, &b)
	fmt.Printf("a= %d,b = %d\n", a, b)
}

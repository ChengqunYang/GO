package main

import "fmt"

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14
	//%T操作变量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

	fmt.Printf("a = %d b = %s c = %c d = %f\n", a, b, c, d)

	//%V自动匹配格式
	fmt.Printf("a = %v b = %v c = %v d = %v", a, b, c, d)

}

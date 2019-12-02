package main

import "fmt"

func jiecheng(a int) int {
	if a == 1 {
		return 1
	} else {
		return a * jiecheng(a-1)
	}

}
func main() {
	var a int
	fmt.Println("请输入要求阶乘的数:")
	fmt.Scan(&a)

	result := jiecheng(a)
	fmt.Println(result)
}

package main

import "fmt"

func main() {
	//自动推导类型
	var a []int
	a = append(a, 34, 6)
	fmt.Println(a)
	//自动推导类型,并初始化
	b := []int{1, 2, 3, 0, 0}
	s := b[1:3:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) // 长度: 3-1
	fmt.Println("cap(s) = ", cap(s)) //容量: 5-1
	//借助make函数初始化,如果不指定最后的那个10(容量)那么容量默认等于长度
	c := make([]int, 5, 10)
	fmt.Println("c = ", c)
	fmt.Println("len(c) = ", len(c))
	fmt.Println("cap(c) = ", cap(c))
}

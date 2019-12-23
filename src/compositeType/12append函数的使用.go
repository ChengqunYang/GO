package main

import (
	"fmt"
)

func main() {
	//append函数,在原切片的末尾添加元素
	s1 := []int{}
	fmt.Printf("len = %d,cap = %d\n", len(s1), cap(s1))
	fmt.Println(s1)
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Printf("len = %d,cap = %d\n", len(s1), cap(s1))
	fmt.Println(s1)

	s2 := []int{2, 3, 4}
	fmt.Println("s2 = ", s2)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Println(s2)

	//如果超过原来的容量,通常二倍的容量扩容
	s := make([]int, 0, 1) //创建一个切片,长度为0,容量为1
	oldcap := cap(s)
	fmt.Println(oldcap)
	for i := 0; i < 9; i++ {
		s = append(s, i)
		fmt.Println("cap = ", cap(s))
	}
}

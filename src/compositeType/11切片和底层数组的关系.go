package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//新切片
	s1 := a[2:5] //从a[2]开始,取三个元素
	s1[1] = 666
	fmt.Println(s1)
	fmt.Println(a)

	s2 := s1[2:7]
	fmt.Println(s2)
	s2[2] = 777
	fmt.Println(a)
}

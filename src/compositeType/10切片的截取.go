package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max] 取下标从low开始的元素,leng-high-low,cap-max-low
	s1 := array[:] //等价于[0:len(arry):len(array)] 不指定容量和长度相等
	fmt.Println(s1)

	fmt.Printf("len = %d,cap = %d\n", len(s1), cap(s1))
	//操作某个元素,和数组的操作方式一样
	data := array[0]
	fmt.Println(data)

	s2 := array[3:6:7]
	fmt.Println(s2)
	fmt.Printf("len = %d,cap = %d\n", len(s2), cap(s2))

	s3 := array[:5] //从零开始取五个元素,容量等于array的长度
	fmt.Println(s3)
	fmt.Printf("len = %d,cap = %d\n", len(s3), cap(s3))

	s4 := array[3:] //从下标3开始(包含),到结尾,容量等于长度
	fmt.Println(s4)
	fmt.Printf("len = %d,cap = %d\n", len(s4), cap(s4))

}

package main

import "fmt"

func main() {
	//数组的定义,元素个数是常量
	var id [10]int
	//操作数组,通过下标
	/*	for i := range id {
			id[i] = i + 1
		}
		for i := 0; i < len(id); i++ {
			fmt.Printf("id[%d] = %d\n", i,id[i] )
		}*/
	for i, data := range id {
		b := 0
		data = b
		b++
		fmt.Printf("id[%d] = %d\n", i, data)
	}

	//部分初始化,没有初始化的元素,自动赋值为0
	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)

	//指定某个元素初始化
	d := [5]int{2: 29, 4: 15}
	fmt.Println(d)

}

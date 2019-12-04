package main

import "fmt"

func main() {
	var id [50]int
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
}

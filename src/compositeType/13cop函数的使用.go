package main

import "fmt"

func main() {
	srcSlic := []int{1, 2}
	dsrSlice := []int{6, 6, 6, 6, 6}
	fmt.Println(srcSlic)
	fmt.Println(dsrSlice)
	copy(dsrSlice, srcSlic)
	fmt.Println("dst = ", dsrSlice)
	//dst =  [1 2 6 6 6]
}

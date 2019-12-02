package main

import "fmt"

func MyFunc09(a, b int) (max, min int) {

	if a > b {
		max = a
		min = b
	} else {
		min = a
		max = b
	}
	return
}
func main() {
	max, min := MyFunc09(10, 20)
	fmt.Println(max, min)
}

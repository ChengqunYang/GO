package main

import "fmt"

func main() {
	//有多少个方括号就是多少维
	var a [3][4]int
	k := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	b := [5]int{1, 2, 4, 7}
	c := [5]int{1, 2, 4, 7}
	d := [5]int{2, 4, 6}
	fmt.Println("a == b", b == c)
	fmt.Println("c ==d", c == d)
	b = d
	fmt.Println("b = ", b)
}

package main

import "fmt"

func main() {
	var a int
	var b float32
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	var (
		e int
		f float64
	)
	e, f = 10, 20.3
	fmt.Println("e = ", e, "f = ", f)
	var d = 19
	fmt.Println(d)

	const c int = 70
	const g float64 = 3.14
	fmt.Println(c, g)

	const (
		i int     = 80
		j float64 = 3.25
	)
	fmt.Println(i, j)
}

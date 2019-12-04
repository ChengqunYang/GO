package main

import "fmt"

func main() {
	var p *int
	fmt.Println("p = ", p)

	//runtime error: invalid memory address or nil pointer dereference
	//*p = 666

	var a int
	p = &a
	*p = 888
	fmt.Println("p = ", a)
}

package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer func() {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}()

	//参数会先传递进去,只是没有执行,在最后执行而已
	defer func(a, b int) {
		fmt.Printf("有参的:a= %d,b=%d\n", a, b)
	}(a, b)

	a = 111
	b = 222
	fmt.Println(a, b)
}

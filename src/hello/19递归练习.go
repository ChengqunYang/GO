package main

import "fmt"

func jia3(n int) int {
	if n==1||n==2||n==3 {
		return 1
	}else {
		return jia3(n-1)+jia3(n-2)+jia3(n-3)
	}
}
func main() {
	var a int
	fmt.Println("请输入项数:")
	fmt.Scanln(&a)
	result :=jia3(a)
	fmt.Printf("第%d项的值为%d.\n",a,result)
	fmt.Println("该项除以3的余数为:",result%3)
}

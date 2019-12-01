package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for { //for后面不写任何东西
		i++
		if i == 5 {
			//break   //跳出循环   可以用于for,switch,break
			continue //跳出本次循环,仅能用于for循环
		}
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

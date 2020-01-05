package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)
	timer.Reset(time.Second)
	<-timer.C
	fmt.Println("时间到")
}

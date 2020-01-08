package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
}

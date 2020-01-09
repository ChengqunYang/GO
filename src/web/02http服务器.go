package main

import (
	"fmt"
	"net/http"
)

func main() {
	//注册处理函数,用户连接,自动调用指定的处理函数
	http.HandleFunc("/", HandConn)

	//监听绑定
	http.ListenAndServe(":9000", nil)

}

//w是用于给客户端回复数据的,req读取客户端发送的数据
func HandConn(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("request.Method = ", request.Method)
	fmt.Println("request.URL = ", request.URL)
	fmt.Println("request.Header = ", request.Header)

	writer.Write([]byte("hello go")) //给客户端回复数据
}

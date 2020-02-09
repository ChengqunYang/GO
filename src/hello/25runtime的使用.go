package main

import (
	"fmt"
	"path"
	"runtime"
)

//runtime.Caller
func main() {
	pc, file, line, ok := runtime.Caller(0) //参数表示显示的行数是调用者还是他本身的位置
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

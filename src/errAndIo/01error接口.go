package main

import (
	"errors"
	"fmt"
)

func main() {
	//两种方式生成error这个接口的实现类的对象
	err1 := fmt.Errorf("%s", "this is normal err1")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is normal err2")
	fmt.Println("err2 = ", err2)
}

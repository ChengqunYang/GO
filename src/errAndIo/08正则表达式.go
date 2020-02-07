package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := " abc azc a7c aac 888 a9c tac"
	//1.解释规则
	reg1 := regexp.MustCompile("a.c") //它会解析正则表达式,如果成功返回解释器
	if reg1 == nil {                  //说明解释失败
		fmt.Println(" regexp err")
		return
	}
	//2.根据规则提前关键信息
	result := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)
}

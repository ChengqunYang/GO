package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 agsgsd 1.23 7. 8.99 lgdadfur 6.66"

	//定义一个正则表达式,+表示匹配多次
	var reg = regexp.MustCompile(`\d+\.\d+`) //提取关键信息
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	//result:=reg.FindAllString(buf,-1)
	result := reg.FindAllStringSubmatch(buf, -1) //会分组
	fmt.Println("result = ", result)
}

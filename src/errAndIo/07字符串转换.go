package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串的转化

	//1.Append系列函数将整数等转化为字符串后,添加到现有的字节数组中
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10) //第二个为要追加的数,第三个数说明要以几进制的方式追加
	fmt.Println(string(str))
	str = strconv.AppendQuote(str, "asdf")
	fmt.Println(string(str))
	fmt.Println("----------")

	//2.把其他类型的转化为字符串
	fmt.Println(strconv.FormatBool(false))
	fmt.Println(strconv.FormatInt(8, 2)) //第二个参数为将这个数转化为几进制形式的字符串
	//f指打印格式,以小数的方式,-1指小数点位数(紧缩模式),64以float64处理
	fmt.Println(strconv.FormatFloat(3.14, 'f', -1, 64))
	fmt.Println("----------")

	//3. 字符串转化为其他类型,如果不能转则会报错误error
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}
	atoi, _ := strconv.Atoi("567")
	fmt.Println(atoi)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串常用操作

	//1.字符串s中是否包含substr,返回bool值
	fmt.Println(strings.Contains("aaabbbcccddd", "ccd"))
	fmt.Println("----------")

	//2.字符串连接,把slice a 通过sep练级诶起来
	s := []string{"aaa", "bbb", "ccc"}
	fmt.Println(strings.Join(s, ","))
	fmt.Println("----------")

	//3.在字符串内查找sep所在的位置,返回位置下标,找不到返回-1
	fmt.Println(strings.Index("csdgfasdg", "f"))
	fmt.Println("----------")

	//4.重复字符串count次,最后返回重复的字符串
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println("----------")

	//5.在s字符串中将old字符串替换为new字符串,n表示替换的次数,小于0表示全部替换
	fmt.Println(strings.Replace("babamama", "ba", "yeyenainai", 5))
	fmt.Println("----------")

	//6.把字符串s按照sep分割,返回slice
	fmt.Println(strings.Split("ba,ba,ma,ma,ye,nai", ","))
	fmt.Println("----------")

	//7.在字符串的头部和尾部去掉指定的字符串
	fmt.Println(strings.Trim("!!!!hehahhea!!", "!!!"))
	fmt.Println("----------")

	//8.去除字符串的空格符,并且按照空格分割返回slice
	fmt.Println(strings.Fields("asdf asdf hgf wer fgh"))
	fmt.Println("----------")

}

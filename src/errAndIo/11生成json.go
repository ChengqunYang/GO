package main

import (
	"encoding/json"
	"fmt"
)

/*json是一种比xml更加轻量级的数据交换格式,在易于人们阅读和编写的同时,也易于程序解析和生成,
尽管json是javaScript的一个子集,但json采用完全独立于编程语言的文本格式,
且表现为键值对集合的文本描述形式,这使他成为了理想的跨平台跨语言的数据交换语言
*/
type IT struct {
	Company  string   `json:"company"`  //二次编码  如果写成string `json:"_"`表示此字段不会被转化为json
	Subjects []string `json:"subjects"` //二次编码,设定转化为json后,对应字段的键首字母为小写
	IsOk     bool     `json:",string"`  //二次编码,表示转化后的json中此字段的类型为string
	Price    float64
}

func main() {
	//开发者可以使用json传递简单的字符串,数字,布尔值,也可以传输一个数组,
	//或者一个更复杂的复合结构,在web开发领域中,json被广泛应用于web服务器和客户端之间的数据通信
	//生成json文本信息
	//1.通过结构体生成json
	var s IT
	s = IT{"itcast", []string{"go", "c++", "python"}, true, 666.666}
	//编码,根据结构体内容生成json
	//buf, err := json.Marshal(s)  //返回一个字节切片
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
	fmt.Println("----------")
	//2.通过map生成json
	//创建一个键为string,值为任意类型的,容量为4的map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subJects"] = []string{"go", "c++", "python"}
	m["isOk"] = true
	m["price"] = 666.666

	//生成json
	result, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println(string(result))
}

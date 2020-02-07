package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"` //二次编码
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:",string"`
	Price    float64
}

func main() {
	jsonbuf := `{
	"company": "itcast",
		"isOk": "true",
		"price": 666.666,
		"subJects": [
		"go",
		"c++",
		"python"
		]
}`
	//1.将json转化为struct
	var temp IT
	if err := json.Unmarshal([]byte(jsonbuf), &temp); err != nil {
		fmt.Println("err = ", err)
	} else {
		//fmt.Println("temp = ",temp)
		fmt.Println("temp = ", temp)
	}
	fmt.Println("--------------")
	//2.将json转化为map
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(jsonbuf), &m); err != nil {
		fmt.Println("err = ", err)
		return
	} else {
		fmt.Println(m)
	}
	fmt.Println("--------------")
	//var str string = m["company"] //错误:因为不确定这个map的值到底是什么类型,所以不能直接赋给一个字符串
	//类型断言
	var str string
	for key, value := range m {
		//fmt.Printf("%v =====%v\n",key,value)
		/*if key == "company" {
			str = value
			fmt.Println("str = ",str)
		}*/
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的类型为string,值为%s\n", key, str)
		case []interface{}:
			fmt.Printf("map[%s]的类型为string,值为%v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的类型为float64,值为%f\n", key, data)
		}
	}
}

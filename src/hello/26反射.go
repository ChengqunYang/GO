package main

import (
	"fmt"
	"reflect"
)

/*
反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下三个。

	1. 基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发panic，那很可能是在代码写完的很长时间之后。
	2. 大量使用反射的代码通常难以理解。
	3. 反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢一到两个数量级
*/
/*
reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。
reflect.Value与原始值之间可以互相转换。
	Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
	Int() int64	将值以 int 类型返回，所有有符号整型均可以此方式返回
	Uint() uint64	将值以 uint 类型返回，所有无符号整型均可以此方式返回
	Float() float64	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
	Bool() bool	将值以 bool 类型返回
	Bytes() []bytes	将值以字节数组 []bytes 类型返回
	String() string	将值以字符串类型返回

isNil()
	func (v Value) IsNil() bool
	IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。

isValid()
	func (v Value) IsValid() bool
	IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
*/
type Cat struct {
}

//通过反射设置变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) //修改的是副本,reflect包也会引发panic
	}
}
func main() {
	var a float32 = 3.14
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	fmt.Println("___________")

	var b float64 = 10.03
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println("___________")

	var c int64 = 100
	fmt.Println(reflect.TypeOf(c))
	vf := reflect.ValueOf(c)
	fmt.Println(vf)
	fmt.Printf("tyep is %s,value is %d\n", reflect.TypeOf(c), vf)
	fmt.Println("___________")

	var d = Cat{}
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(d).Name())
	fmt.Println(reflect.TypeOf(d).Kind())
	fmt.Println(reflect.ValueOf(d))
	fmt.Println("___________")

	reflectSetValue(&c)
	fmt.Println(c)
	fmt.Println("____________")

	//IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	var e *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(e).IsNil())
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	fmt.Println("______________")
}

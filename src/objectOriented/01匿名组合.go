package main

type Person struct {
	name string
	sex byte
	age int
}
type Student struct {
	Person  //匿名字段,那么默认Student就包含了Person的所有字段
	id int
	addr string
}
func main() {

}

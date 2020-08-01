package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)
// ES insert data demo
type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Married bool`json:"married"`
}
func main() {
	// 1.初始化一个连接,得到一个client
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es success")
	p1 := Student{
		Name:    "aaa",
		Age:     21,
		Married: false,
	}
	put1, err := client.Index().Index("student").Type("go").BodyJson(p1).Do(context.Background())
	if err != nil {
		// Handle err
		panic(err)
	}
	fmt.Printf("Indexed student %s to index %s,type %s\n",put1.Id,put1.Index,put1.Type)
}

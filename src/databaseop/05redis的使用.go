package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     "127.0.0.1:6379",
		Password: "",
	})
	_, err = redisdb.Ping().Result()
	return
}
func redisExample() {
	err := redisdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := redisdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := redisdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}
func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed,err:%v\n", err)
		return
	}
	redisExample()
}

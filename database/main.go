package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-redis/redis"
)

type People struct {
	Name string
	Age  int
}

var client *redis.Client

func ConnectRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleSetKey(key string, value interface{}) {
	v, _ := json.Marshal(value)

	/* func (c cmdable) Set(key string, value interface{}, expiration time.Duration) *StatusCmd */
	err := client.Set(key, v, 0).Err()
	if err != nil {
		panic(err)
	}
}

func ExampleGetKey(key string) {
	val, err := client.Get(key).Result()

	/* Không tồn tại key */
	if err == redis.Nil {
		fmt.Println("'", key, "' does not exist")
		return
	}

	/* Xảy ra lỗi */
	if err != nil {
		panic(err)
	}

	/* Lấy được key */
	j := People{}
	json.Unmarshal([]byte(val), &j)

	fmt.Println(key, j)
	fmt.Println("type", reflect.TypeOf(client))
}

func main() {
	key := "people"
	value := People{Name: "name", Age: 12}

	ConnectRedis()
	ExampleSetKey(key, value)
	ExampleGetKey(key)
}

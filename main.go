package main
import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

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
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func ExampleSetKey(key string, value interface{}, expire time.Duration) {
	v, _ := json.Marshal(value)

	err := client.Set(key, v, expire).Err()
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
	fmt.Println("type", reflect.TypeOf(j))
}

func main() {
	key := "people"
	value := People{Name: "name", Age: 12}

	ConnectRedis()
	ExampleSetKey(key, value, 0)
	ExampleGetKey(key)
}

package redis

import (
	"fmt"

	red "github.com/go-redis/redis"
)

var rdb *red.Client

func initClient() (err error) {
	rdb = red.NewClient(&red.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func redisExampleSet() {
	err := rdb.Set("go-test-set", 100, 0).Err()
	if err != nil {
		fmt.Printf("set failed. err:%v\n", err)
		return
	}
}

func redisExampleGet() {
	val, err := rdb.Get("go-test-set").Result()
	if err != nil {
		fmt.Printf("get failed. err:%v\n", err)
		return
	}
	fmt.Println("value", val)
}

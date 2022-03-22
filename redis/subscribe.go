package main

import (
	"log"

	"github.com/go-redis/redis/v8"
)

func subscriber(key string) interface{} {
	rdb := connectionRedis()

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println("key does not exist: ", key)
	} else if err != nil {
		panic(err)
	}

	defer rdb.Close()

	return val
}

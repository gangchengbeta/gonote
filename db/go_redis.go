package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	redisClient *redis.Client
	_           context.Context
)

func RedisCRUD() {
	initRedis()
	//readString()
	//writeString()
	//deleteRedis()
	//writeList()
	//readAllList()
	//writeHash()
	readHash()
}

func readHash() {
	key := "hash"
	value, err := redisClient.HGet(key, "id").Result()
	handleRedisError(err)
	fmt.Println(value)
	redisClient.Del(key)
}

func writeHash() {
	key := "hash"
	value, err := redisClient.HSet(key, "id", "123").Result()
	handleRedisError(err)
	fmt.Println(value)
}

func readAllList() {
	key := "list"
	value, err := redisClient.LRange(key, 0, -1).Result()
	handleRedisError(err)
	fmt.Println(value)
}

func writeList() {
	key := "list"
	value := []interface{}{1, 2, 3, "傻逼张堡垒"}
	// 写入
	err := redisClient.LPush(key, value...).Err()
	//redisClient.Expire(key, 10*time.Second)
	handleRedisError(err)
}

func deleteRedis() {
	key := "name"
	err := redisClient.Del(key).Err()
	handleRedisError(err)
}

func readString() {
	key := "name"
	value, err := redisClient.Get(key).Result()
	handleRedisError(err)
	fmt.Println(value)
}
func initRedis() {
	// 不影响主程序
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "112.126.71.240:6379",
		Password: "dockerredis", //
		DB:       0,
	})
	//ctx = context.TODO()
	_, err := redisClient.Ping().Result()
	handleRedisError(err)
	fmt.Println("redis连接成功！")
}

func writeString() {
	key := "name"
	value := "coder4j"
	err := redisClient.Set(key, value, 10*time.Second).Err()
	// 也可以调用Expire方法设置过期时间
	//redisClient.Expire(key, 10*time.Second)
	handleRedisError(err)
}

// key不存在时 不至于结束程序
func handleRedisError(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("key does not exist\n")
			return
		}
		panic(err)
	}
}

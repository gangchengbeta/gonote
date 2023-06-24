package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	RedisClient *redis.Client
	_           context.Context
)

func RedisCRUD() {
	//readString()
	//writeString()
	//deleteRedis()
	//writeList()
	//readAllList()
	writeHash()
	//readHash()
}

func readHash() {
	key := "student:123"
	value, err := RedisClient.HGet(key, "id").Result()
	HandleRedisError(err)
	fmt.Println(value)
	RedisClient.Del(key)
}

func writeHash() {
	key := "student:"
	filed := make(map[string]interface{})
	filed["id"] = "123"
	filed["Name"] = "coder4j"
	filed["Age"] = 18
	filed["Height"] = 180.04
	_, err := RedisClient.HMSet(key+"123", filed).Result()
	HandleRedisError(err)
}

func readAllList() {
	key := "list"
	value, err := RedisClient.LRange(key, 0, -1).Result()
	HandleRedisError(err)
	fmt.Println(value)
}

func writeList() {
	key := "list"
	value := []interface{}{1, 2, 3, "傻逼张堡垒"}
	// 写入
	err := RedisClient.LPush(key, value...).Err()
	//RedisClient.Expire(key, 10*time.Second)
	HandleRedisError(err)
}

func deleteRedis() {
	key := "name"
	err := RedisClient.Del(key).Err()
	HandleRedisError(err)
}

func readString() {
	key := "name"
	value, err := RedisClient.Get(key).Result()
	HandleRedisError(err)
	fmt.Println(value)
}
func init() {
	// 不影响主程序
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "112.126.71.240:6379",
		Password: "dockerredis", //
		DB:       0,
	})
	//ctx = context.TODO()
	_, err := RedisClient.Ping().Result()
	HandleRedisError(err)
	fmt.Println("redis连接成功！")
}

func writeString() {
	key := "name"
	value := "coder4j"
	err := RedisClient.Set(key, value, 10*time.Second).Err()
	// 也可以调用Expire方法设置过期时间
	//RedisClient.Expire(key, 10*time.Second)
	HandleRedisError(err)
}

// key不存在时 不至于结束程序
func HandleRedisError(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("key does not exist\n")
			return
		}
		panic(err)
	}
}

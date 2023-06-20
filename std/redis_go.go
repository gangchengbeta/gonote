package std

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	// 不影响主程序
	defer func() {
		recover()
	}()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "112.126.71.240:6379",
		Password: "dockerredis", //
		DB:       0,
	})
	err := initRedis()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("redis连接成功！")
}

// redis 初始化
func initRedis() (err error) {
	_, err = RedisClient.Ping().Result()
	return
}

// RedisBasic go 连接redis
func RedisBasic() {
	s, _ := RedisClient.Get("login:token:u7fpffqgduw17y3eh5zomltaunfockaz").Result()
	fmt.Println(s)
	defer func(redisclient *redis.Client) {
		err := redisclient.Close()
		if err != nil {
			panic(err)
		}
	}(RedisClient)
}

package std

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisclient *redis.Client

func initRedis() (err error) {
	redisclient = redis.NewClient(&redis.Options{
		Addr:     "112.126.71.240:6379",
		Password: "dockerredis", //
		DB:       0,
	})
	_, err = redisclient.Ping().Result()
	return
}

// go 连接redis

func RedisBasic() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed! err : %v\n", err)
		return
	}
	fmt.Println("redis连接成功！")
	s, err := redisclient.Get("login:token:u7fpffqgduw17y3eh5zomltaunfockaz").Result()
	fmt.Println(s)

	defer func(redisclient *redis.Client) {
		err := redisclient.Close()
		if err != nil {
			panic(err)
		}
	}(redisclient)
}

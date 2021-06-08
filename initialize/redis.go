package initialize

import (
	"context"
	"fmt"
	"yasi_audio/global"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		print("失败")
		fmt.Println(err)
	} else {
		fmt.Println("成功", pong)
		global.GVA_REDIS = client
	}
}

package initialize

import (
	"context"
	"yasi_audio/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
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
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}

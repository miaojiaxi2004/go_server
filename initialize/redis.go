package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/miaojiaxi2004/go_server/global"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	fmt.Println(pong)
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
		return nil
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}

package initialize

import (
	"context"
	"fmt"

	"github.com/quangdvn/go-ec/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	cache := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // No password set
		DB:       r.Database, // Use default DB
		PoolSize: r.PoolSize, //
	})
	_, err := cache.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("redis init error", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("redis connected")
	global.Cache = cache
	redisExample()
}

func redisExample() {
	err := global.Cache.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Println("error redis:", zap.Error(err))
		return
	}
	value, err := global.Cache.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("error redis:", zap.Error(err))
		return
	}
	global.Logger.Info("redis example::::", zap.String("value", value))
}

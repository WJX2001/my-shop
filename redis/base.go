package redis

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
)

var (
	RdsConn *redis.Client
	once    sync.Once
)

func init() {
	once.Do(func() {
		redisConfig, _ := beego.AppConfig.GetSection("redis")
		db_num, _ := strconv.Atoi(redisConfig["db_index"])
		RdsConn = redis.NewClient(&redis.Options{
			Addr:     redisConfig["address"],
			Password: redisConfig["password"],
			DB:       db_num,
		})
		ctx := context.Background()
		_, err := RdsConn.Ping(ctx).Result()
		if err != nil {
			logs.Info("connect redis fail", err)
			panic(err)
		}
	})
}

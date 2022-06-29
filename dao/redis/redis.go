package redis

import (
	"bluebell/settings"
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	if _, err = rdb.Ping().Result(); err != nil {
		return
	}

	return
}

func Close() {
	_ = rdb.Close()
}

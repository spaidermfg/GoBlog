package setting

import (
	"github.com/go-redis/redis/v9"
)


func InitRedisFunc(cfg *RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Host,
		Password: cfg.Password,
		DB: cfg.Db,
	})

	return rdb
}


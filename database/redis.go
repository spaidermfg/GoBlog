package database

import (
	"GoBlog/setting"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"time"
)

var rdb *redis.Pool

func InitRedis(cfg *setting.RedisConfig) {
	rdb = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port )
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				zap.L().Error("init redis fail...")
				return nil, err
			}
			zap.L().Debug("init redis success...")
			_ ,er := conn.Do("set","k9","999")
			if er != nil {
				zap.L().Error("redis set key fail")
			}
			val, ers := conn.Do("get", "k9")
			if ers == nil{
				fmt.Println(val)
			}

			return conn, nil
		},

		MaxActive:cfg.MaxActive,  //最大空闲连接数
		MaxIdle: cfg.MaxIdle,     //数据库最大链接数

		IdleTimeout: time.Duration(cfg.IdleTimeout),   //最大空闲时间


	}
}


//关闭连接池
func CloseRdb() {
	_ = rdb.Close()
}
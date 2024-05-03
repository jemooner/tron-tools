package redis

import (
	"github.com/go-redis/redis"
	"tron-tools/config"
)

// RedisDB Redis的DB对象
var RedisDB *redis.Client

func NewRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Host,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Database,
	})

	defer func() {
		if r := recover(); r != nil {
			logger.Error("Redis connection error,", r)
		}
	}()
	_, err := RedisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
	logger.Info("Redis connection successful!")
}

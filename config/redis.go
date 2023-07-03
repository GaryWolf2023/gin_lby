package config

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdb *redis.Client
var ndration = 30 * 24 * 60 * 60 * time.Second // 30天过期

type RedisClient struct{}

func InitRedis() (*RedisClient, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: viper.GetString("redis.password"), // 没有密码，默认值
		DB:       viper.GetInt("redis.DB"),          // 默认DB 0
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any) error {
	return rdb.Set(context.Background(), key, value, ndration).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return rdb.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return rdb.Del(context.Background(), key...).Err()
}

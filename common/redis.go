package common

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var ctx = context.Background()
var RedisClient *redis.Client

func InitRedis() *redis.Client {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	AddInfo("redis init successfully!")

	return RedisClient
}

// 以下是具体的redis操作封装 //

func Set(key string, value interface{}) error {
	return RedisClient.Set(ctx, key, value, 0).Err()
}

func Get(key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

func Del(key string) error {
	return RedisClient.Del(ctx, key).Err()
}

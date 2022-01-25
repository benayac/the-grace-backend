package helper

import (
	"github.com/go-redis/redis"
	"thegrace/pkg"
	"time"
)

var redisClient *redis.Client

func getRedisClient() *redis.Client {
	if redisClient != nil {
		return redisClient
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     pkg.Conf.RedisHost,
		Password: "",
		DB:       0,
	})
	return redisClient
}

func SetValue(key string, value string, expiry time.Duration) error {
	client := getRedisClient()
	err := client.Set(key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetValue(key string) (string, error) {
	client := getRedisClient()
	value, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

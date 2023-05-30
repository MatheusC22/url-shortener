package services

import (
	"context"
	"goAPI/configs"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisService() *RedisService {
	var new_ctx = context.Background()
	var new_client = createClient()
	return &RedisService{ctx: new_ctx, client: new_client}
}

func createClient() *redis.Client {
	var conf = configs.GetRedis()
	client := redis.NewClient(&redis.Options{
		Addr:       conf.Addr,
		Password:   conf.Password,
		DB:         conf.Db,
		MaxConnAge: 2 * time.Second,
	})
	return client
}

func (c *RedisService) Set(key string, val string) error {
	err := c.client.Set(c.ctx, key, val, 86400*time.Second).Err() //EXPIRATION DATE 1 DAY 86400
	if err == nil {
		return nil
	}
	return err
}

func (c *RedisService) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

func (c *RedisService) Del(key string) error {
	return c.client.Del(c.ctx, key).Err()
}

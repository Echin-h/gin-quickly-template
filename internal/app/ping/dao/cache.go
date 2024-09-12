package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type pingCache struct {
	*redis.Client
}

func (u *pingCache) InitCache(rds *redis.Client) error {
	u.Client = rds
	return u.Ping()
}

func (u *pingCache) Ping() error {
	_, err := u.Client.Ping(context.Background()).Result()
	return err
}

func (u *pingCache) Set(key string, value interface{}, expiration int64) error {
	return u.Client.Set(context.Background(), key, value, 0).Err()
}

func (u *pingCache) Get(key string) (string, error) {
	return u.Client.Get(context.Background(), key).Result()
}

func (u *pingCache) Del(key string) error {
	return u.Client.Del(context.Background(), key).Err()
}

func (u *pingCache) Exists(key string) (int64, error) {
	return u.Client.Exists(context.Background(), key).Result()
}

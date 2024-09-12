package redis

import (
	"context"
	"fmt"
	"gin-quickly-template/config"
	"gin-quickly-template/pkg/colorful"
	"github.com/redis/go-redis/v9"
)

func InitCache() *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().Database.Redis.Addr,
		Password: config.GetConfig().Database.Redis.Password,
		DB:       config.GetConfig().Database.Redis.DB,
	})

	p := rds.Ping(context.Background())
	if p.Err() != nil {
		fmt.Println(colorful.Red("redis connect failed, err: " + p.Err().Error()))
		return nil
	}

	fmt.Println(colorful.Green("redis connect success"))
	return rds
}

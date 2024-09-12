package dao

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	PingOrm   = &pingOrm{}
	PingCache = &pingCache{}
)

func Init(db *gorm.DB, rds *redis.Client) error {
	err := PingOrm.InitPG(db)
	if err != nil {
		return err
	}

	err = PingCache.InitCache(rds)
	if err != nil {
		return err
	}

	return nil
}

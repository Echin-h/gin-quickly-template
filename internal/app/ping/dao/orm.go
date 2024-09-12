package dao

import (
	"context"
	"gin-quickly-template/internal/app/ping/model"
	"gorm.io/gorm"
)

type pingOrm struct {
	*gorm.DB
}

func (u *pingOrm) InitPG(db *gorm.DB) error {
	u.DB = db
	return db.AutoMigrate(&model.Ping{})
}

func (u *pingOrm) CreatePing(ping model.Ping) error {
	return u.DB.WithContext(context.Background()).Create(&ping).Error
}

func (u *pingOrm) GetPingList() ([]model.Ping, error) {
	var pings []model.Ping
	err := u.DB.WithContext(context.Background()).Find(&pings).Error
	return pings, err
}

func (u *pingOrm) GetPingByID(id int) (model.Ping, error) {
	var pin model.Ping
	err := u.DB.WithContext(context.Background()).First(&pin, id).Error
	return pin, err
}

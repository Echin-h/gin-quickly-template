package model

type Ping struct {
	ID      int    `gorm:"primaryKey"`
	Message string `gorm:"type:varchar(255);not null"`
}

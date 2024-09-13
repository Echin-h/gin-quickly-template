package pgsql

import (
	"fmt"
	"gin-quickly-template/config"
	"gin-quickly-template/pkg/colorful"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitPostgres() *gorm.DB {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.GetConfig().Database.Postgres.Host,
		config.GetConfig().Database.Postgres.Port,
		config.GetConfig().Database.Postgres.Username,
		config.GetConfig().Database.Postgres.Password,
		config.GetConfig().Database.Postgres.DBName,
		config.GetConfig().Database.Postgres.SSLMode,
		config.GetConfig().Database.Postgres.TimeZone,
	)

	ormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info), // it can be change
	}

	db, err := gorm.Open(postgres.Open(dns), ormConfig)
	if err != nil {
		fmt.Println(colorful.Red("postgres connect failed, err: " + err.Error()))
		return nil
	}

	//if config.GetConfig().OTel.Enable {
	//	err = db.Use(otelgorm.NewPlugin())
	//	errorx.PanicOnErr(err)
	//}

	fmt.Println(colorful.Green("postgres connect success"))
	return db
}

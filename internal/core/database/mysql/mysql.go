package mysql

import (
	"fmt"
	"gin-quickly-template/config"
	"gin-quickly-template/pkg/colorful"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		config.GetConfig().Database.Mysql.Username,
		config.GetConfig().Database.Mysql.Password,
		config.GetConfig().Database.Mysql.Host,
		config.GetConfig().Database.Mysql.Port,
		config.GetConfig().Database.Mysql.DBName,
		config.GetConfig().Database.Mysql.Charset,
		config.GetConfig().Database.Mysql.ParseTime,
		config.GetConfig().Database.Mysql.Loc,
	)

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 使用单数表名
		Logger:         logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		fmt.Println(colorful.Red("mysql connect failed, err: " + err.Error()))
		return nil
	}

	//if config.GetConfig().OTel.Enable {
	//	err = db.Use(otelgorm.NewPlugin())
	//	errorx.PanicOnErr(err)
	//}

	fmt.Println(colorful.Green("mysql connect success"))

	return db
}

package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Gorm = &_Gorm{}

type _Gorm struct {
	conn *gorm.DB
}

func (db *_Gorm) GetInstance() *gorm.DB {
	return db.conn
}

func (db *_Gorm) InitGorm() (conn *gorm.DB) {
	var conf = mysql.Config{
		DriverName:        Config.GetString("DB.DRIVER"),
		DSN:               Config.GetString("DB.DSN"),
		DefaultStringSize: 191,
	}
	conn, err := gorm.Open(
		mysql.Open(conf.DSN),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second,   // 慢 SQL 阈值
					LogLevel:      logger.Silent, // Log level
					Colorful:      true,          // 禁用彩色打印
				},
			),
		})
	if err != nil {
		panic(err.Error() + "\n failed to connect database \n")
	}
	db.conn = conn
	return
}

package utils

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var MySQL = &_MySQL{}

type _MySQL struct {
	conn *gorm.DB
}

func (db *_MySQL) GetInstance() *gorm.DB {
	return db.conn
}

func (db *_MySQL) InitMySQL() (conn *gorm.DB) {
	var conf = mysql.Config{
		DriverName:                Config.GetString("db.driver"),
		DSN:                       Config.GetString("db.dsn"),
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	sqlDB, err := sql.Open(conf.DriverName, conf.DSN)
	if err != nil {
		panic(err.Error() + "\n failed to connect database \n")
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间。
	//sqlDB.SetConnMaxLifetime(time.Hour)
	//defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		panic(err.Error() + "\n database ping failed \n")
	}

	conn, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
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

//func ConnectMySQL() *sql.DB {
//	DriverName := "mysql"
//	Dsn := "root:secret@tcp(mysql:3306)/demo?charset=utf8&parseTime=True&loc=Local"
//	db, err := sql.Open(DriverName, Dsn)
//	if err != nil {
//		panic(err.Error())
//	}
//	//defer db.Close()
//
//	//设置数据库最大连接数
//	db.SetConnMaxLifetime(100)
//	//设置上数据库最大闲置连接数
//	db.SetMaxIdleConns(10)
//
//	err = db.Ping()
//	if err != nil {
//		panic(err.Error())
//	}
//	return db
//}

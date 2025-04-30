package di

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitMysql() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{

		//本地地址
		//DSN: "root:root@tcp(127.0.0.1:3306)/campus?charset=utf8mb4&parseTime=True&loc=Local",
		DSN: "root:0429@tcp(127.0.0.1:3306)/hotel?charset=utf8mb4&parseTime=True&loc=Local",

		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数
		},
		// 打印sql语句
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("mysql init fault, error is" + err.Error())
	}

	// 配置链接池
	Db, _ := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	Db.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Db.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Db.SetConnMaxLifetime(time.Hour)

}

// Conn
//
//	@desc: 获取链接
//	@param ctx context.Context
//	@return *gorm.DB
func Conn(ctx context.Context) *gorm.DB {
	return DB.WithContext(ctx)
}

package initialize

import (
	"fmt"
	"github.com/lopo1/mall/auth/global"
	"github.com/lopo1/mall/auth/model"
	"go.uber.org/zap"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB(){
	c := global.ServerConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,         // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	MysqlTables(global.DB)
}

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		fmt.Println("register table failed", err)
		zap.S().Error("register table failed", err)
		os.Exit(0)
	}
	zap.S().Info("register table success")
}

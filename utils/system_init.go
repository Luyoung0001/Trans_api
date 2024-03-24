package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./config/") //带绝对路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
}

// 初始化数据库MySQL

func InitMySQL() {

	// 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	// 连接数据库
	var err error
	fmt.Println(viper.GetString("mysql.dns"))
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}
	// 测试数据库连接
	fmt.Println("Database connection successful!")
}

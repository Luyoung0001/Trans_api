package utils

import (
	"fmt"
	"github.com/spf13/viper"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Initconfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("../config") //带绝对路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
}

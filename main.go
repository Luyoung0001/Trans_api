package main

import (
	"github.com/spf13/viper"
	"log"
	"trans_api/modules"
	"trans_api/router"
	"trans_api/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	bind()

	r := router.Router()
	err := r.Run(viper.GetString("port.server"))
	if err != nil {
		return
	}

}

func bind() {
	// 自动迁移模型
	err := utils.DB.AutoMigrate(&modules.WordTranslation{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	err = utils.DB.AutoMigrate(&modules.Words{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	err = utils.DB.AutoMigrate(&modules.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	err = utils.DB.AutoMigrate(&modules.RelationBookWord{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}

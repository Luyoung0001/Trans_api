package main

import (
	"github.com/spf13/viper"
	"log"
	"trans_api/modules"
	"trans_api/router"
	"trans_api/utils"
)

// @title           trans_api
// @version         1.0
// @description     This is a celler server
// @termsOfService  http://swagger.io/terms/
// @contact.name   luyoung
// @contact.email  luyoung0001@gmail.com
// @host      localhost:9000
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
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

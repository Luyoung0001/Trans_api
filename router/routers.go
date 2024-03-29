package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"trans_api/docs"
	"trans_api/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// restFUL api

	r.GET("/word/", service.Eng2Chi)     // 查询英语
	r.GET("/english/", service.SearchEn) // 查询英语单词的一切信息
	r.GET("/chinese/", service.Chi2Eng)  // 查询汉字
	return r
}

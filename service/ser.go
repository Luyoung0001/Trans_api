package service

import (
	"github.com/gin-gonic/gin"
	"trans_api/modules"
)

// SearchResponse 查询返回结果结构体
type SearchResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Eng2Chi 查询单词的中文翻译
// @Summary 查询单词的中文翻译
// @Description 根据传入的英语单词查询对应的中文翻译
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param english FormData string false "英语单词"
// @Success 200 {object} SearchResponse "成功返回"
// @Router /word [get]
func Eng2Chi(c *gin.Context) {
	// 查询数据库
	english := c.Request.FormValue("english")
	data := modules.SearchEng(english)
	c.JSON(200, SearchResponse{
		Code:    0,
		Message: "查询单词成功!",
		Data:    data,
	})
}

// Chi2Eng 查询中文翻译的英语单词
// @Summary 查询中文翻译的英语单词
// @Description 根据传入的中文单词查询对应的英语翻译
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param chinese query string false "中文单词"
// @Success 200 {object} gin.H{"code": 0, "message": "查询单词成功!", "data": modules.SearchChiResponse} "成功返回"
// @Router /word [get]
//func Chi2Eng(c *gin.Context) {
//	// 查询数据库
//	chinese := c.Query("chinese")
//	data := modules.SearchChi(chinese)
//	c.JSON(200, gin.H{
//		"code":    0,
//		"message": "查询单词成功!",
//		"data":    data,
//	})
//}

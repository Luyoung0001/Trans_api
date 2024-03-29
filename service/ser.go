package service

import (
	"github.com/gin-gonic/gin"
	"trans_api/modules"
)

// SearchResponse 定义查询返回结果的结构体
type SearchResponse struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 消息
	Data    interface{} `json:"data"`    // 数据
}

// Eng2Chi 用于查询英语单词的中文翻译
// @Summary 查询英语单词的中文翻译
// @Description 根据传入的英语单词查询对应的中文翻译
// @Tags 用户模块
// @Produce json
// @Param english query string true "英语单词"
// @Success 200 {object} SearchResponse "查询成功"
// @Failure 400 {object} SearchResponse "查询失败"
// @Router /word [get]
func Eng2Chi(c *gin.Context) {
	english := c.Query("english") // 使用 Query 方法获取 URL 中的参数
	data := modules.SearchEng(english)
	// 如果查询成功，返回 400 状态码
	if data == nil {
		c.JSON(400, SearchResponse{
			Code:    0,
			Message: "查询单词失败!",
			Data:    nil,
		})
		return
	}
	// 如果查询失败，返回其他状态码
	c.JSON(200, SearchResponse{
		Code:    1,
		Message: "查询单词成功!",
		Data:    data,
	})
}

// Chi2Eng 用于查询中文翻译的英语单词
// @Summary 查询中文翻译的英语单词
// @Description 根据传入的中文单词查询对应的英语翻译
// @Tags 用户模块
// @Produce json
// @Param chinese query string true "中文单词"
// @Success 200 {object} SearchResponse "查询成功"
// @Failure 400 {object} SearchResponse "查询失败"
// @Router /chinese [get]
func Chi2Eng(c *gin.Context) {
	chinese := c.Query("chinese") // 使用 Query 方法获取 URL 中的参数
	data := modules.SearchChi(chinese)
	// 如果查询成功，返回 400 状态码
	if data == nil {
		c.JSON(400, SearchResponse{
			Code:    0,
			Message: "查询单词失败!",
			Data:    nil,
		})
		return
	}

	// 如果查询失败，返回其他状态码
	c.JSON(200, SearchResponse{
		Code:    1,
		Message: "查询单词成功!",
		Data:    data,
	})
}

// SearchEn 用于查询英文单词的一切信息
// @Summary 查询英文单词的一切信息
// @Description 根据传入的英文单词查询对应的英语翻译
// @Tags 用户模块
// @Produce json
// @Param english query string true "英文单词"
// @Success 200 {object} SearchResponse "查询成功"
// @Failure 400 {object} SearchResponse "查询失败"
// @Router /english [get]
func SearchEn(c *gin.Context) {
	english := c.Query("english") // 使用 Query 方法获取 URL 中的参数
	data := modules.SearchWord(english)
	// 如果查询成功，返回 400 状态码
	if data == nil {
		c.JSON(400, SearchResponse{
			Code:    0,
			Message: "查询单词失败!",
			Data:    nil,
		})
		return
	}

	// 如果查询失败，返回其他状态码
	c.JSON(200, SearchResponse{
		Code:    1,
		Message: "查询单词成功!",
		Data:    data,
	})
}

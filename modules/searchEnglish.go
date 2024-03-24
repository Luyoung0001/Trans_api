package modules

import (
	"fmt"
	"trans_api/utils"
)
// 简单查询单词的意思
func SearchEng(english string) WordTranslation {
	data := WordTranslation{}
	fmt.Print(english)
	utils.DB.Where("word = ?", english).First(&data)
	fmt.Print(data)
	return data
}
// 查询单词的一切信息
func SearchWord(english string) Word {
	data := Word{}
	fmt.Print(english)
	utils.DB.Where("vc_vocabulary = ?", english).First(&data)
	fmt.Print(data)
	return data
}
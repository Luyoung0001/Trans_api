package modules

import (
	"trans_api/utils"
)

// 简单查询单词的意思

func SearchEng(english string) WordTranslation {
	data := WordTranslation{}
	utils.DB.Where("word = ?", english).First(&data)
	return data
}

// 查询单词的一切信息

func SearchWord(english string) Words {
	data := Words{}
	utils.DB.Where("vc_vocabulary = ?", english).First(&data)
	return data
}

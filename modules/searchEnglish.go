package modules

import (
	"trans_api/utils"
)

// 简单查询单词的意思

func SearchEng(english string) *WordTranslation {
	data := WordTranslation{}
	utils.DB.Where("word = ?", english).First(&data)
	// 如果 id 为 0,就是查不到
	if data.ID == 0 {
		return nil
	}
	return &data
}

// 查询单词的一切信息

func SearchWord(english string) *Words {
	data := Words{}
	utils.DB.Where("vc_vocabulary = ?", english).First(&data)
	// 如果 id 为 0,就是查不到
	if data.ID == "" {
		return nil
	}
	return &data
}

package modules

import (
	"trans_api/utils"
)

// 简单查询中文的意思

func SearchChi(chinese string) WordTranslation {
	data := WordTranslation{}
	utils.DB.Where("translation LIKE ?", "%"+chinese+"%").First(&data)
	return data
}

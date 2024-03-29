package modules

import (
	"trans_api/utils"
)

// 简单查询中文的意思

func SearchChi(chinese string) *WordTranslation {
	data := WordTranslation{}
	utils.DB.Where("translation LIKE ?", "%"+chinese+"%").First(&data)
	// 如果 id 为 0,就是查不到
	if data.ID == 0 {
		return nil
	}
	return &data
}

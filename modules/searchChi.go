package modules
import (
	"fmt"
	"trans_api/utils"
)

// 简单查询中文的意思
func SearchChi(chinese string) WordTranslation {
	data := WordTranslation{}
	fmt.Print(chinese)
	utils.DB.Where("translation LIKE ?", "%" + chinese + "%").First(&data)
	fmt.Print(data)
	return data
}
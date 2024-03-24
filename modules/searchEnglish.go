package modules

import (
	"fmt"
	"trans_api/utils"
)

func SearchEng(english string) WordTranslation {
	data := WordTranslation{}
	fmt.Print(english)
	utils.DB.Where("word = ?", english).First(&data)
	fmt.Print(data)
	return data
}

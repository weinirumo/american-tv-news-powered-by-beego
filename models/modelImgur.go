package models

//http://jinzhu.me/gorm/ gorm 文档

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Imgur struct {
	gorm.Model
	imgur_id               string
	title                  string
	title_tanslation       string
	description            string
	description_tanslation string
	keywords               string
	Images                 []Image
}

func FetchAllImgurCached() (imgurs []Imgur) {

	if x, found := CacheManager.Get(CK_Imgur_ALL); found {
		buffer := x.([]byte)
		json.Unmarshal(buffer, &imgurs)
	} else {
		Gorm.Preload("Images").Find(&imgurs)
		buffer, _ := json.Marshal(imgurs)
		CacheManager.Set(CK_Imgur_ALL,buffer, C_EXPIRE_TIME_HOUR_01)
	}
	return
}

func (imgur *Imgur) AfterFind() (err error) {
	//装换excerpt
	return
}

package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"my-ganji-app/common"
	"my-ganji-app/types"
)

type ImageFile struct {
	BaseModel
	Id      int64  `json:"id"`
	Url     string `orm:"unique;size(256);index" json:"url"`
	ImgType int8   `json:"img_type"` // 0:用户头像 1:商品评论图片
}

func (imf *ImageFile) TableName() string {
	return common.TableName("user_image")
}

func (imf *ImageFile) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(imf)
}

func (imf *ImageFile) GetImageById(id int64) (*ImageFile, int, error) {
	var image ImageFile
	err := image.Query().Filter("Id", id).One(&image)
	if err != nil {
		return nil, 100, err
	}
	return &image, types.ReturnSuccess, nil
}

package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/pkg/errors"
	"my-ganji-app/common"
	"my-ganji-app/types"
)

type GoodsImage struct {
	BaseModel
	Id       int64  `orm:"pk;column(id);auto" json:"id" form:"id"`
	GoodsId  int64  `json:"goods_id"`
	Image    string `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"image"` // 商品图片
	IsDispay int8   `orm:"default(1)" json:"is_dispay"`                                                   // 0 不显示 1 显示
}

func (g *GoodsImage) TableName() string {
	return common.TableName("goods_image")
}

func GetGoodsImgList(goods_id int64) ([]*GoodsImage, int, error) {
	var goods_img_list []*GoodsImage
	if _, err := orm.NewOrm().QueryTable(GoodsImage{}).Filter("GoodsId", goods_id).All(&goods_img_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return goods_img_list, types.ReturnSuccess, nil
}

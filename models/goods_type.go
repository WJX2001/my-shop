package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/pkg/errors"
	"my-ganji-app/common"
	"my-ganji-app/types"
)

type GoodsType struct {
	BaseModel
	Id       int64  `orm:"column(id);auto;size(11)" json:"id"`
	GoodsId  int64  `json:"goods_id"`                   // 商品ID
	TypeKey  string `orm:"size(512)" json:"type_key"`   // 属性名称：如颜色，输入商品的人可以自定义
	TypeVale string `orm:"size(1024)" json:"type_vale"` // 属性文字：入库数据格式 ["白色", "蓝色", "黄色"]
	IsDispay int8   `orm:"default(1)" json:"is_dispay"` // 0 不显示 1 显示
}

func (g *GoodsType) TableName() string {
	return common.TableName("goods_type")
}

func GetGoodsTypeList(goods_id int64) ([]*GoodsType, int64, error) {
	var type_list []*GoodsType
	if _, err := orm.NewOrm().QueryTable(GoodsType{}).Filter("GoodsId", goods_id).All(&type_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return type_list, types.ReturnSuccess, nil
}

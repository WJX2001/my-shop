package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/pkg/errors"
	"my-ganji-app/common"
	"my-ganji-app/types"
)

type GoodsCar struct {
	BaseModel
	Id         int64   `orm:"column(id);auto;size(11)" json:"id" form:"id"`
	GoodsId    int64   `orm:"default(1)" json:"goods_id"`                                                               // 商品ID
	MerchantId int64   `json:"merchant_id"`                                                                             // 商品所属商家ID
	Logo       string  `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"` // 商品LOGO
	AddresId   int64   `orm:"default(1)" json:"addres_id"`                                                              // 地址ID
	GoodsTypes string  `orm:"size(512)" json:"goods_types"`                                                             // 商品属性
	GoodsTitle string  `orm:"size(64)" json:"goods_title"`                                                              // 商品标题
	GoodsName  string  `orm:"size(512);index" json:"goods_name" form:"goods_name"`                                      // 产品名称
	UserId     int64   `orm:"size(64);index" json:"user_id"`                                                            // 购买用户
	BuyNums    int64   `orm:"default(0)" json:"buy_nums"`                                                               // 购买数量
	PayAmount  float64 `orm:"default(0);digits(22);decimals(8)" json:"pay_amount"`                                      // 支付金额
}

func (c *GoodsCar) TableName() string {
	return common.TableName("goods_car")
}

func (c *GoodsCar) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *GoodsCar) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func GetGoodsCarDetailByGoodsId(user_id, goods_id int64) (*GoodsCar, int, error) {
	goods_car := GoodsCar{}
	if err := orm.NewOrm().QueryTable(GoodsCar{}).
		Filter("UserId", user_id).
		Filter("goods_id", goods_id).One(&goods_car); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &goods_car, types.ReturnSuccess, nil
}

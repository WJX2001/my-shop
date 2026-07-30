package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"my-ganji-app/common"
)

type Goods struct {
	BaseModel
	Id             int     `orm:"column(id);auto" json:"id" form:"id"`
	GoodsCatId     int64   `json:"goods_cat_id"`                                                                            // 商品所属一级分类ID
	GoodsLastCatId int64   `json:"goods_level_cat_id"`                                                                      // 商品所属最后一级分类ID
	GoodsMark      string  `orm:"size(512);index" json:"goods_mark"`                                                        // 商品备注
	Serveice       string  `orm:"size(512);index" json:"serveice"`                                                          // 服务说明
	CalcWay        int8    `orm:"default(0);index" json:"calc_way"`                                                         // 0:按件计量 1:按近计量
	MerchantId     int64   `json:"merchant_id"`                                                                             // 商品所属商家ID
	Title          string  `orm:"size(512);index" json:"title"`                                                             // 商品标题
	Logo           string  `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"` // 商品封面
	TotalAmount    int64   `orm:"default(150000)" json:"total_amount" form:"total_amount"`                                  // 商品总量
	LeftAmount     int64   `orm:"default(150000)" json:"left_amount" form:"left_amount"`                                    // 剩余商品总量
	GoodsPrice     float64 `orm:"default(1);digits(22);decimals(8)" json:"goods_price"`                                     // 商品价格
	GoodsDisPrice  float64 `orm:"default(1);digits(22);decimals(8)" json:"goods_discount_price"`                            // 商品折扣价格
	GoodsIntegral  float64 `orm:"default(1);digits(22);decimals(8)" json:"goods_integral"`                                  // 购买需要的积分数量
	SendIntegral   float64 `orm:"default(1);digits(22);decimals(8)" json:"send_integral"`                                   // 购买商品赠送积分
	GoodsName      string  `orm:"size(512);index" json:"goods_name" form:"goods_name"`                                      // 产品名称
	GoodsParams    string  `orm:"type(text)" json:"goods_params" form:"goods_params"`                                       // 产品参数
	GoodsDetail    string  `orm:"type(text)" json:"goods_detail" form:"goods_detail"`                                       // 产品详细介绍
	Discount       float64 `orm:"default(0);index" json:"discount"`                                                         // 折扣，取值 0.1-9.9；0代表不打折
	Sale           int8    `orm:"default(0);index" json:"sale" form:"sale"`                                                 // 0:上架 1:下架
	IsDisplay      int8    `orm:"default(0);index" json:"is_display" form:"is_display"`                                     // 0:首页不展示, 1:首页展示
	SellNums       int64   `orm:"default(0);index" json:"sell_nums"`                                                        // 售出数量
	IsHot          int8    `orm:"default(0);index" json:"is_hot"`                                                           // 0:非爆款产品 1:爆款产品
	IsDiscount     int8    `orm:"default(0);index" json:"is_discount"`                                                      // 0:不打折，1:打折活动产品
	IsIgSend       int8    `orm:"default(0);index" json:"is_ig_send"`                                                       // 0:正常， 1:赠送积分
	IsGroup        int8    `orm:"default(0);index" json:"is_group"`                                                         // 0:非拼购产品 1:拼购产品
	GroupNumber    int64   `orm:"default(100);index" json:"group_number"`                                                   // 助力人数
	IsIntegral     int8    `orm:"default(0);index" json:"is_integral"`                                                      // 0:非积分兑换产品 1:积分兑换产品
	LeftTime       int64   `orm:"default(0);index" json:"left_time"`                                                        // 限时产品剩余时间
	IsLimitTime    int8    `orm:"default(0);index" json:"is_limit_time"`                                                    // 0:不是限时产品 1:是限时
}

func (g Goods) TableName() string {
	return common.TableName("goods")
}

func GetMerchantGoodsNums(merChant_id int64) int64 {
	total, err := orm.NewOrm().QueryTable(Goods{}).Filter("MerchantId", merChant_id).Count()
	if err != nil {
		return 0
	}
	return total
}

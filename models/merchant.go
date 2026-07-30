package models

import (
	ado "github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/client/orm"
	"github.com/pkg/errors"
	"my-ganji-app/common"
	"my-ganji-app/types"
)

type Merchant struct {
	BaseModel
	Id             int64   `json:"id" form:"id"`
	Logo           string  `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"` // 商家 Logo
	MerchantName   string  `orm:"size(512);index" json:"merchant_name"`                                                     // 商家名称
	MerchantIntro  string  `orm:"size(512);index" json:"merchant_intro"`                                                    // 商家简介
	MerchantDetail string  `orm:"type(text)" json:"merchant_detail"`                                                        // 商家详情
	ContactUser    string  `orm:"size(128);index" json:"contact_user"`                                                      // 商家联系人
	Phone          string  `orm:"size(64);index" json:"phone"`                                                              // 商家联系电话
	WeChat         string  `orm:"size(64);index" json:"we_chat"`                                                            // 商家联系微信
	Address        string  `orm:"size(512);index" json:"address"`                                                           // 店铺地址
	GoodsNum       int64   `json:"goods_num"`                                                                               // 商品总数
	MerchantWay    int8    `orm:"default(0);index" json:"merchant_way"`                                                     // 0:自营商家； 1:认证商家  2:普通商家
	SettlePercent  float64 `orm:"default(0);digits(22);decimals(8)" json:"settle_percent"`                                  // 结算比例
	ShopLevel      int8    `json:"shop_level"`                                                                              // 店铺等级
	ShopServer     int8    `json:"shop_server"`                                                                             // 店铺服务
}

func (m *Merchant) TableName() string {
	return common.TableName("merchant")
}

func GetMerchantList(page, pageSize int, merct_name string, address string) ([]*Merchant, int64, error) {
	offset := (page - 1) * pageSize
	merchant_list := make([]*Merchant, 0)

	// SetCond 需要 client/orm.Condition，不能用 adapter 的 NewCondition
	cond := orm.NewCondition()
	query := ado.NewOrm().QueryTable(new(Merchant))
	if merct_name != "" {
		cond = cond.Or("MerchantName__contains", merct_name)
	}
	if address != "" {
		cond = cond.Or("Address__contains", address)
	}
	if merct_name != "" || address != "" {
		query = query.SetCond(cond)
	}

	total, err := query.Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = query.OrderBy("-GoodsNum").Limit(pageSize, offset).All(&merchant_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return merchant_list, total, nil
}

func GetMerchantDetail(id int64) (*Merchant, int, error) {
	var merchant Merchant
	if err := orm.NewOrm().QueryTable(Merchant{}).Filter("Id", id).One(&merchant); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &merchant, types.ReturnSuccess, nil
}

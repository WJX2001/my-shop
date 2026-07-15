package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"my-ganji-app/common"
)

type UserWallet struct {
	BaseModel
	Id          int64   `json:"id"`
	UserId      int64   `orm:"index" json:"user_id"`
	AssetName   string  `orm:"size(128);index" json:"asset_name"` // 资产名称
	TotalAmount float64 `orm:"default(150);digits(22);decimals(8)" json:"total_amount"`
}

func (u *UserWallet) TableName() string {
	return common.TableName("user_wallet")
}

func (u *UserWallet) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

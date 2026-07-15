package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"my-ganji-app/common"
)

type UserInfo struct {
	BaseModel
	Id         int64  `json:"id"`
	UserId     int64  `orm:"index" json:"user_id"`
	RealName   string `orm:"default(ganji);size(15);index" json:"real_name"`
	IdCard     string `orm:"default(000000000000000000);size(18)"`                             // 身份证号码
	CardImgPos string `orm:"size(150);default(/static/upload/default/user-default-60x60.png)"` // 身份证正面
	CardImgNeg string `orm:"size(150);default(/static/upload/default/user-default-60x60.png)"` // 身份证反面
	WeChat     string `orm:"default(ganji);size(15);index" json:"we_chat"`
	QQ         string `orm:"default(ganji);size(15);index" json:"qq"`
	Sex        int8   `orm:"default(0);index"` // 0: 男, 1: 女 3:未知
}

func (u *UserInfo) TableName() string {
	return common.TableName("user_info")
}

func (u *UserInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

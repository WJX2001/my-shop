package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"my-ganji-app/common"
	"time"
)

type UserCoupon struct {
	BaseModel
	Id          int64      `json:"id"`
	UserId      int64      `orm:"index" json:"user_id"`
	CouponName  string     `orm:"size(128);index" json:"coupon_name"` // 优惠券名称
	IsUsed      int8       `orm:"default(0)" json:"is_used"`          // 0：未使用；1: 已经使用
	TotalAmount float64    `orm:"default(150);digits(22);decimals(8)" json:"total_amount"`
	StartTime   *time.Time `orm:"type(datetime);index" json:"start_time"`
	EndTime     *time.Time `orm:"type(datetime);index" json:"end_time"`
	IsInvalid   int8       `orm:"default(0)" json:"is_invalid"` // 0 未过期； 1:已经过期
}

func (u *UserCoupon) TableName() string {
	return common.TableName("user_coupon")
}

func (u *UserCoupon) SearchField() []string {
	return []string{"conpon_name"}
}

func (u *UserCoupon) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

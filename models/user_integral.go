package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/pkg/errors"
	"my-ganji-app/common"
)

type UserIntegral struct {
	BaseModel
	Id           int64   `json:"id"`
	UserId       int64   `orm:"size(64);index" json:"user_id"`
	IntegralName string  `orm:"size(128);index" json:"integral_name"`
	TotalIg      float64 `orm:"default(0);digits(22);decimals(8)" json:"total_ig"` // 总的积分
	UsedIg       float64 `orm:"default(0);digits(22);decimals(8)" json:"UsedIg"`   // 已赠送LSDT积分
	TodayIg      float64 `orm:"default(0);digits(22);decimals(8)" json:"today_ig"` // 今日已赠送LSDT积分
}

func (u *UserIntegral) TableName() string {
	return common.TableName("user_integral")
}

func (u *UserIntegral) SearchField() []string {
	return []string{"integral_name"}
}

func (u *UserIntegral) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (u *UserIntegral) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

// UpdateDb 这里将 db orm.Ormer 传入而不是直接用各自的 orm.NewOrm()
// 主要原因是 <支持事务> 如果方法内部自己 NewOrm() 会新建会话，事务失效
func (u *UserIntegral) UpdateDb(db orm.Ormer) error {
	if _, err := db.Update(u); err != nil {
		return err
	}
	return nil
}

func (u *UserIntegral) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func (u *UserIntegral) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}

func (u *UserIntegral) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func GetIntegralByUserId(userIdParam int64) (*UserIntegral, error) {
	var userIg UserIntegral
	err := userIg.Query().Filter("UserId", userIdParam).Limit(1).One(&userIg)
	if err != nil {
		return nil, err
	}
	return &userIg, nil
}

func GetIgByUserId(db orm.Ormer, user_id int64) (*UserIntegral, error) {
	var userIg UserIntegral
	err := db.QueryTable(UserIntegral{}).Filter("UserId", user_id).Limit(1).One(&userIg)
	if err != nil {
		return nil, errors.New("err in GetIgByUserId")
	}
	return &userIg, nil
}

package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/pkg/errors"
	"my-ganji-app/common"
	"my-ganji-app/types"
	"time"
)

type UserAddress struct {
	Id        int64     `json:"id"`
	UserId    int64     `orm:"default(150000)" json:"user_id"`
	UserName  string    `orm:"size(128);index" json:"user_name"` // 收件名字
	Phone     string    `orm:"size(32);index" json:"phone"`      // 手机号
	Address   string    `orm:"size(512);index" json:"address"`   // 地址
	IsSet     int8      `orm:"index" json:"is_set"`              // 0: 正常，1: 默认地址
	Status    int8      `orm:"index" json:"status"`              // 0: 正常，1: 删除
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime);index" json:"updated_at"`
}

func (u *UserAddress) SearchField() []string {
	return []string{"username", "phone", "user_id"}
}

func (u *UserAddress) TableName() string {
	return common.TableName("user_address")
}

func (u *UserAddress) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(u); err != nil {
		return err, 0
	}
	return nil, id
}

func GetUserAddressDefault(user_id int64) (*UserAddress, int, error) {
	address := UserAddress{}
	if err := orm.NewOrm().QueryTable(UserAddress{}).Filter("UserId", user_id).Filter("IsSet", 1).Limit(1).One(&address); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &address, types.ReturnSuccess, nil
}

func (u *UserAddress) GetAddressById() (*UserAddress, int64, string) {
	var address UserAddress
	if err := orm.NewOrm().QueryTable(u.TableName()).Filter("Id", u.Id).One(&address); err != nil {
		return nil, types.SystemDbErr, "数据库查询失败，请联系客服处理"
	}
	return &address, types.ReturnSuccess, ""
}

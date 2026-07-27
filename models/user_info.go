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

func (u *UserInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}
func (u *UserInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *UserInfo) GetUserInfoByUserId(user_id int64) (UserInfo, error) {
	var user_info UserInfo
	err := user_info.Query().Filter("UserId", user_id).Limit(1).One(&user_info)
	if err != nil {
		return user_info, err
	}
	return user_info, nil
}

//// UpdateUserInfo 修改用户信息
//func (u *UserInfo) UpdateUserInfo(user_auth type_user.UserAuthCheck) (success bool, code int, err error) {
//	var user_info UserInfo
//	if err := u.Query().Filter("UserId", u.UserId).One(&user_info); err != nil {
//		return false, types.UserIsNotExist, errors.New("用户不存在")
//	}
//	user_info.RealName = user_auth.RealName
//	user_info.IdCard = user_auth.IdCard
//	var imgfile ImageFile
//	if user_auth.IdCardNegImgId > 0 {
//		img, _, err := imgfile.GetImageById(user_auth.IdCardNegImgId)
//		if err != nil {
//			return false, types.UserAuthError, errors.New("用户实名认证失败，照片上传失败")
//		}
//		user_info.CardImgNeg = img.Url
//	}
//	if user_auth.IdCardPosImgId > 0 {
//		img, _, err := imgfile.GetImageById(user_auth.IdCardPosImgId)
//		if err != nil {
//			return false, types.UserAuthError, errors.New("用户实名认证失败，照片上传失败")
//		}
//		user_info.CardImgPos = img.Url
//	}
//	err = user_info.Update()
//	if err != nil {
//		return false, types.UserAuthError, errors.New("用户实名认证失败，数据库更新失败")
//	}
//	return true, types.ReturnSuccess, nil
//}

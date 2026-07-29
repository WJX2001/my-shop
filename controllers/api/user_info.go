package api

import (
	"encoding/json"
	"my-ganji-app/models"
	"my-ganji-app/types"
	type_user "my-ganji-app/types/user"
)

type UserInfoController struct {
	BaseController
}

func (uic *UserInfoController) GetUserInfo() {
	user, ok := uic.CurrentUser()
	if !ok {
		return
	}
	var userInfo models.UserInfo
	uinf, err := userInfo.GetUserInfoByUserId(user.Id)
	if err != nil {
		uic.Data["json"] = RetResource(false, types.UserIsNotExist, nil, "用户不存在，请联系客服处理")
		uic.ServeJSON()
		return
	}

	var user_integral float64
	user_ig, _ := models.GetIntegralByUserId(user.Id)
	if user_ig != nil {
		user_integral = user_ig.TotalIg
	} else {
		user_integral = 0
	}

	var user_wallet_cny float64
	user_w, _ := models.GetWalletByUserId(user.Id)
	if user_w != nil {
		user_wallet_cny = user_w.TotalAmount
	} else {
		user_wallet_cny = 0
	}
	user_infos := type_user.UserInfoRet{
		UserId:      user.Id,
		Token:       user.Token,
		UserName:    user.UserName,
		IgAmount:    user_integral,
		CnyAmount:   user_wallet_cny,
		Phone:       user.Phone,
		Eamil:       user.Email,
		Sex:         uinf.Sex,
		IsAuth:      user.IsAuth,
		MemberLevel: user.MemberLevel,
		InviteCode:  user.MyInviteCode,
		Avator:      user.Avatar,
		RealName:    uinf.RealName,
		WeiChat:     uinf.WeChat,
		QQ:          uinf.QQ,
	}

	uic.Data["json"] = RetResource(true, types.ReturnSuccess, user_infos, "获取用户信息成功")
	uic.ServeJSON()
	return
}

// UpdateUserInfo 修改用户信息
func (uic *UserInfoController) UpdateUserInfo() {
	user, ok := uic.CurrentUser()
	if !ok {
		return
	}
	var user_info type_user.UpdateUserInfoCheck
	if err := json.Unmarshal(uic.Ctx.Input.RequestBody, &user_info); err != nil {
		uic.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		uic.ServeJSON()
		return
	}
	success, code, err := models.UpdateUserInfo(user.Id, user_info)
	if success {
		uic.Data["json"] = RetResource(true, types.ReturnSuccess, user_info, "修改用户信息成功")
	} else {
		uic.Data["json"] = RetResource(false, code, err, err.Error())
	}
	uic.ServeJSON()
	return
}

// IsAuth 实名认证
func (uic *UserInfoController) IsAuth() {
	user, ok := uic.CurrentUser()
	if !ok {
		return
	}
	var userInfo models.UserInfo
	uinf, _ := userInfo.GetUserInfoByUserId(user.Id)
	auth_data := type_user.UserAuthRet{
		Id:         user.Id,
		Phone:      user.Phone,
		UserName:   user.UserName,
		RealName:   uinf.RealName,
		IdCard:     uinf.IdCard,
		CardImgPos: uinf.CardImgPos,
		CardImgNeg: uinf.CardImgNeg,
		IsAuth:     user.IsAuth,
	}
	uic.Data["json"] = RetResource(true, types.ReturnSuccess, auth_data, "获取用户认证信息成功")
	uic.ServeJSON()
	return
}

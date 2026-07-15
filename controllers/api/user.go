package api

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"my-ganji-app/models"
	"my-ganji-app/types"
	typeUser "my-ganji-app/types/user"
	"strings"
)

type UserController struct {
	beego.Controller
}

func (uc *UserController) GetUserInfo() {
	bearerToken := uc.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		uc.Data["json"] = RetResource(false, 4000, nil, "您还没有登陆，请登陆")
		err := uc.ServeJSON()
		if err != nil {
			return
		}
		return
	}

	token := strings.TrimPrefix(bearerToken, "Bearer ")
	userToken, err := models.GetUserByToken(token)
	if err != nil {
		uc.Data["json"] = RetResource(false, 4000, nil, "您还有没有登陆，请登陆")
		err := uc.ServeJSON()
		if err != nil {
			return
		}
		return
	}

	userInfo := make(map[string]string)
	userInfo["user_name"] = userToken.UserName
	userInfo["email"] = userToken.Email
	userInfo["userToken"] = userToken.Token
	uc.Data["json"] = RetResource(true, 2000, userInfo, "获取我的邀请码成功")
	err = uc.ServeJSON()
	if err != nil {
		return
	}
	return
}

// UserRegister @Title UserRegister
// @Description 用户注册 UserRegister
// @Success 200 status bool, data interface{}, msg string
// @router /register [post]
func (uc *UserController) UserRegister() {
	ctx := uc.Ctx.Request.Context()
	registerParam := typeUser.UserRegisterCheck{}
	if err := json.Unmarshal(uc.Ctx.Input.RequestBody, &registerParam); err != nil {
		uc.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		uc.ServeJSON()
		return
	} else {
		if code, err := registerParam.UserRegisterCheckParamValidate(ctx); err != nil {
			uc.Data["json"] = RetResource(false, code, nil, err.Error())
			uc.ServeJSON()
			return
		}

		success, code, err := models.RegisterByPhoneOrEmail(ctx, registerParam)
		if success {
			uc.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "")
		} else {
			msg := "注册失败"
			if err != nil {
				msg = err.Error()
			}
			uc.Data["json"] = RetResource(false, code, nil, msg)
		}
	}

	uc.ServeJSON()
	return
}

func (uc *UserController) UserLogin() {
	ctx := uc.Ctx.Request.Context()
	loginParam := typeUser.UserLoginCheck{}
	if err := json.Unmarshal(uc.Ctx.Input.RequestBody, &loginParam); err != nil {
		uc.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		uc.ServeJSON()
		return
	} else {
		if code, err := loginParam.UserLoginCheckParamValidate(ctx); err != nil {
			uc.Data["json"] = RetResource(false, code, nil, err.Error())
			uc.ServeJSON()
			return
		}
		_, user_data, code, err := models.LoginByPhoneOrEmail(loginParam)
		if code == types.ReturnSuccess {
			uc.Data["json"] = RetResource(true, types.ReturnSuccess, user_data, "登陆成功")
		} else {
			uc.Data["json"] = RetResource(false, code, nil, err.Error())
		}
		uc.ServeJSON()
		return
	}
}

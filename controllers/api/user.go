package api

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"my-ganji-app/common"
	"my-ganji-app/common/utils"
	"my-ganji-app/models"
	rds_conn "my-ganji-app/redis"
	"my-ganji-app/types"
	type_user "my-ganji-app/types/user"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	beego.Controller
}

// SendPhoneCode @Title SendPhoneCode
// @Description 发送手机号验证码 SendPhoneCode
// @Success 200 status bool, data interface{}, msg string
// @router /send_phone_code [post]
func (uc *UserController) SendPhoneCode() {
	ctx := uc.Ctx.Request.Context()
	phone_number := type_user.PhoneNumberCheck{}
	if err := json.Unmarshal(uc.Ctx.Input.RequestBody, &phone_number); err != nil {
		uc.Data["json"] = RetResource(false, types.InvalidFormatError, err.Error(), "无效的参数格式，请联系客服处理")
		uc.ServeJSON()
		return
	} else {
		if code, err := phone_number.PhoneNumberParamValidate(); err != nil {
			uc.Data["json"] = RetResource(false, code, nil, err.Error())
			uc.ServeJSON()
			return
		}
		verify_code, _ := strconv.Atoi(common.GenValidateCode(6))
		// 打印生成的验证码，控制台可见
		logs.Info("当前手机号：%s，生成验证码：%d", phone_number.Phone, verify_code)

		rds_conn.RdsConn.Del(ctx, phone_number.Phone)
		rds_conn.RdsConn.Set(ctx, phone_number.Phone, fmt.Sprintf("%d", verify_code), time.Duration(1000)*time.Second).Err()
		utils.SendMessageCode(phone_number.Phone, verify_code)
		uc.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "发送手机号验证码成功")
		uc.ServeJSON()
		return
	}
}

func (uc *UserController) PhoneCodeCheck() {
	ctx := uc.Ctx.Request.Context()
	phone_code_check := type_user.PhoneCodeCheck{}
	if err := json.Unmarshal(uc.Ctx.Input.RequestBody, &phone_code_check); err != nil {
		uc.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		uc.ServeJSON()
		return
	} else {
		if code, err := phone_code_check.ReqPhoneCodeCheckParamValidate(ctx); err != nil {
			uc.Data["json"] = RetResource(false, code, nil, err.Error())
			uc.ServeJSON()
			return
		}
	}
	uc.Data["json"] = RetResource(false, types.ReturnSuccess, nil, "手机号验证校验成功")
	uc.ServeJSON()
	return
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
	registerParam := type_user.UserRegisterCheck{}
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
	loginParam := type_user.UserLoginCheck{}
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

package user

import (
	"context"
	"github.com/pkg/errors"
	rds_conn "my-ganji-app/redis"
	"my-ganji-app/types"
	"regexp"
)

const (
	PhoneNumRule = "^(1[3|4|5|6|7|8|9][0-9]\\d{4,8})$"
	EmailPattern = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
)

type UserRegisterCheck struct {
	VerifyWay      int8   `json:"verify_way"` // 1: 手机号码验证； 2：邮箱验证
	PhoneEmail     string `json:"phone_email"`
	PhoneEmailCode string `json:"phone_email_code"`
	Password1      string `json:"password1"`
	Password2      string `json:"password2"`
	InviteCode     string `json:"invite_code"`
}

func (urc UserRegisterCheck) UserRegisterCheckParamValidate(ctx context.Context) (int, error) {
	if urc.VerifyWay == 1 { // 手机号码验证
		if urc.PhoneEmailCode == "" {
			return types.PhoneVerifyCodeEmptyError, errors.New("手机号验证码为空")
		}
		result, _ := regexp.MatchString(PhoneNumRule, urc.PhoneEmail)
		if !result {
			return types.PhoneFormatError, errors.New("手机号码格式不正确")
		}
		phone_code := rds_conn.RdsConn.Get(ctx, urc.PhoneEmail).Val()

		if phone_code != urc.PhoneEmailCode {
			return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
		}
	} else if urc.VerifyWay == 2 {
		result, _ := regexp.MatchString(EmailPattern, urc.PhoneEmail)
		if !result {
			return types.EmailFormatError, errors.New("邮箱格式不正确")
		}
		email_code := rds_conn.RdsConn.Get(ctx, urc.PhoneEmail).Val()
		if email_code != urc.PhoneEmailCode {
			return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
		}
	} else {
		return types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
	if urc.Password1 == "" || urc.Password2 == "" {
		return types.PasswordIsEmpty, errors.New("输入的密码不能为空")
	}
	if urc.Password1 != urc.Password2 {
		return types.TwicePasswordNotEqual, errors.New("两次输入的密码不一样")
	}
	return types.ReturnSuccess, nil
}

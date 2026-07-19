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

type PhoneNumberCheck struct {
	Phone string `json:"phone"`
}

func (pnc PhoneNumberCheck) PhoneNumberParamValidate() (int, error) {
	result, _ := regexp.MatchString(PhoneNumRule, pnc.Phone)
	if !result {
		return types.PhoneFormatError, errors.New("手机号码格式不正确")
	}
	return types.ReturnSuccess, nil
}

type PhoneCodeCheck struct {
	PhoneNumberCheck
	PhoneCode string `json:"phone_code"`
}

func (pcc PhoneCodeCheck) ReqPhoneCodeCheckParamValidate(ctx context.Context) (int, error) {
	code, err := pcc.PhoneNumberParamValidate()
	if err != nil {
		return code, err
	}
	if pcc.PhoneCode == "" {
		return types.PhoneVerifyCodeEmptyError, errors.New("手机号码验证码为空")
	}
	phone_code := rds_conn.RdsConn.Get(ctx, pcc.Phone).Val()
	if phone_code != pcc.PhoneCode {
		return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
	}
	return types.ReturnSuccess, nil
}

type PhoneRegisterCheck struct {
	PhoneNumberCheck
	LoginRegister int8 `json:"login_register"` // 1: 登陆 2: 注册
}

func (prc PhoneRegisterCheck) PhoneRegisterCheckParamValidate() (int, error) {
	code, err := prc.PhoneNumberParamValidate()
	if err != nil {
		return code, err
	}
	if prc.LoginRegister != 1 && prc.LoginRegister != 2 {
		return types.NoThisLoginRegisterWay, errors.New("没有这种验证方式，请选择 1 或者 2； 1 表示登陆，2表示注册")
	}
	return types.ReturnSuccess, nil
}

type EmailNumberCheck struct {
	Email string `json:"email"`
}

func (enc EmailNumberCheck) EmailNumberCheckParamValidate() (int, error) {
	result, _ := regexp.MatchString(EmailPattern, enc.Email)
	if !result {
		return types.EmailFormatError, errors.New("邮箱格式不正确")
	}
	return types.ReturnSuccess, nil
}

type EmailCodeCheck struct {
	EmailNumberCheck
	EmailCode string `json:"email_code"`
}

func (ecc EmailCodeCheck) EmailCodeCheckParamValidate(ctx context.Context) (int, error) {
	code, err := ecc.EmailNumberCheckParamValidate()
	if err != nil {
		return code, err
	}
	if ecc.EmailCode == "" {
		return types.EmailVerifyCodeEmptyError, errors.New("邮箱验证码为空")
	}
	email_code := rds_conn.RdsConn.Get(ctx, ecc.Email).Val()
	if email_code != ecc.EmailCode {
		return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
	}
	return types.ReturnSuccess, nil
}

type EmailRegisterCheck struct {
	EmailNumberCheck
	LoginRegister int8 `json:"login_register"` // 1: 登陆 2: 注册
}

func (erc EmailRegisterCheck) EmailRegisterCheckParamValidate() (int, error) {
	code, err := erc.EmailNumberCheckParamValidate()
	if err != nil {
		return code, err
	}
	if erc.LoginRegister != 1 && erc.LoginRegister != 2 {
		return types.NoThisLoginRegisterWay, errors.New("没有这种验证方式，请选择 1 或者 2; 1:表示登陆，2:表示注册")
	}
	return types.ReturnSuccess, nil
}

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

type UserLoginCheck struct {
	VerifyWay      int8   `json:"verify_way"`
	PhoneEmail     string `json:"phone_email"`
	PhoneEmailCode string `json:"phone_email_code"`
	Password       string `json:"password"`
}

func (ulc UserLoginCheck) UserLoginCheckParamValidate(ctx context.Context) (int, error) {
	if ulc.VerifyWay == 1 { // 手机号码验证
		if ulc.PhoneEmailCode == "" {
			return types.PhoneVerifyCodeEmptyError, errors.New("手机号验证码为空")
		}
		result, _ := regexp.MatchString(PhoneNumRule, ulc.PhoneEmail)
		if !result {
			return types.PhoneFormatError, errors.New("手机号码格式不正确")
		}
		phoneCode := rds_conn.RdsConn.Get(ctx, ulc.PhoneEmail).Val()
		if phoneCode != ulc.PhoneEmailCode {
			return types.PhoneVerifyCodeError, errors.New("手机验证码不正确")
		}
		if ulc.Password == "" {
			return types.PasswordIsEmpty, errors.New("输入的密码不能为空")
		}
	} else if ulc.VerifyWay == 2 { // 邮箱验证
		result, _ := regexp.MatchString(EmailPattern, ulc.PhoneEmail)
		if !result {
			return types.EmailFormatError, errors.New("邮箱格式不正确")
		}
		emailCode := rds_conn.RdsConn.Get(ctx, ulc.PhoneEmail).Val()
		if emailCode != ulc.PhoneEmailCode {
			return types.EmailVerifyCodeError, errors.New("邮箱验证码错误")
		}
		if ulc.Password == "" {
			return types.PasswordIsEmpty, errors.New("输入的密码不能为空")
		}
	} else {
		return types.InvalidVerifyWay, errors.New("无效的验证方式")
	}
	return types.ReturnSuccess, nil
}

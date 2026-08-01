package address

import (
	"github.com/pkg/errors"
	"my-ganji-app/types"
	"my-ganji-app/types/user"
	"regexp"
)

type UserAddressAddCheck struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	IsSet    int8   `json:"is_set"`
}

func (u UserAddressAddCheck) UserAddressAddCheckParamValidate() (int, error) {
	if u.UserId <= 0 {
		return types.UserIsNotExist, errors.New("用户不存在, 请联系客服处理")
	}
	if u.UserName == "" {
		return types.UserIsNotExist, errors.New("用户名为空，请务必填写")
	}
	if u.Phone == "" {
		return types.PhoneEmptyError, errors.New("手机号码不能为空，请务必填写")
	}
	result, _ := regexp.MatchString(user.PhoneNumRule, u.Phone)
	if !result {
		return types.PhoneFormatError, errors.New("手机号码格式不正确")
	}
	if u.Address == "" {
		return types.AddressIsEmpty, errors.New("地址不能为空，请务必填写")
	}
	return types.ReturnSuccess, nil
}

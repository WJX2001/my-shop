package api

import (
	"encoding/json"
	"my-ganji-app/models"
	"my-ganji-app/types"
	type_address "my-ganji-app/types/address"
)

type UserAddressController struct {
	BaseController
}

func (c *UserAddressController) AddAddress() {
	requestUser, ok := c.CurrentUser()
	if !ok {
		return
	}
	var add_address type_address.UserAddressAddCheck
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &add_address); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		c.ServeJSON()
		return
	}
	if code, err := add_address.UserAddressAddCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, err, err.Error())
		c.ServeJSON()
		return
	}
	if requestUser.Id != add_address.UserId {
		c.Data["json"] = RetResource(false, types.UserIsNotExist, nil, "Token 和 用户不匹配，拒绝添加地址")
		c.ServeJSON()
		return
	}
	address := models.UserAddress{
		UserId:   add_address.UserId,
		UserName: add_address.UserName,
		Phone:    add_address.Phone,
		Address:  add_address.Address,
		IsSet:    add_address.IsSet,
		Status:   0,
	}
	if err, id := address.Insert(); err != nil {
		c.Data["json"] = RetResource(false, types.CreateAddressFail, nil, "创建地址失败，请联系客服处理")
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = RetResource(true, types.ReturnSuccess, map[string]interface{}{"id": id}, "添加地址成功")
		c.ServeJSON()
		return
	}
}

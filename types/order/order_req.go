package order

import (
	"github.com/pkg/errors"
	"my-ganji-app/types"
)

type CreateOrderCheck struct {
	GoodsId    int64  `json:"goods_id"`
	AddressId  int64  `json:"address_id"`
	UserId     int64  `json:"user_id"`
	BuyNums    int64  `json:"buy_nums"`
	PayWay     int8   `json:"pay_way"` // 0:积分兑换，1:账户余额支付，2:微信支付；3:支付宝支付
	GoodsTypes string `json:"goods_types"`
}

func (c CreateOrderCheck) CreateOrderCheckParamValidate() (int, error) {
	if c.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品ID小于等于0")
	}
	if c.AddressId <= 0 {
		return types.ParamLessZero, errors.New("您没有选择地址，或者您还没有添加地址，请去选择或者添加")
	}
	if c.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于0")
	}
	if c.BuyNums <= 0 {
		return types.ParamLessZero, errors.New("购买数量小于等于 0")
	}
	if c.PayWay < 0 || c.PayWay > 3 {
		return types.InvalidVerifyWay, errors.New("无效的付款方式")
	}
	return types.ReturnSuccess, nil
}

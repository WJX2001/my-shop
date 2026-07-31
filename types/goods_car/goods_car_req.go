package goods_car

import (
	"github.com/pkg/errors"
	"my-ganji-app/types"
)

type AddGoodCarCheck struct {
	GoodsId    int64   `json:"goods_id"`
	UserId     int64   `json:"user_id"`
	AddressId  int64   `json:"address_id"`
	BuyNums    int64   `json:"buy_nums"`
	PayAmount  float64 `json:"pay_amount"`
	GoodsTypes string  `json:"goods_types"` // 商品属性
	IsDis      int8    `json:"is_dis"`      // 1: 非打折商品 2: 打折商品
}

func (a AddGoodCarCheck) AddGoodCarCheckParamValidate() (int, error) {
	if a.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品 ID 不能小于等于 0 ")
	}
	if a.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户 ID 不能小于等于 0")
	}
	if a.AddressId <= 0 {
		return types.ParamLessZero, errors.New("您没有选择地址，或者您还没有添加地址，请去选择或者添加")
	}
	if a.BuyNums <= 0 {
		return types.ParamLessZero, errors.New("购买数量不能小于等于 0")
	}
	if a.PayAmount <= 0 {
		return types.ParamLessZero, errors.New("支付金额不能小于等于 0")
	}
	return types.ReturnSuccess, nil
}

type GoodCarListCheck struct {
	types.PageSizeData
}

func (g GoodCarListCheck) GoodCarListCheckParamValidate() (int, error) {
	code, err := g.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	return types.ReturnSuccess, nil
}

package goods_car

import (
	"github.com/pkg/errors"
	"my-ganji-app/types"
)

type AddGoodCarCheck struct {
	GoodsId    int64  `json:"goods_id"`
	UserId     int64  `json:"user_id"`
	AddressId  int64  `json:"address_id"`
	BuyNums    int64  `json:"buy_nums"`
	GoodsTypes string `json:"goods_types"` // 商品属性
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
	return types.ReturnSuccess, nil
}

type EditGoodCarCheck struct {
	GoodsId    int64 `json:"goods_id"`
	GoodsCarId int64 `json:"goods_car_id"`
	UserId     int64 `json:"user_id"`
	BuyNums    int64 `json:"buy_nums"`
}

func (e EditGoodCarCheck) EditGoodCarCheckParamValidate() (int, error) {
	if e.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品 ID 不能小于等于 0")
	}
	if e.GoodsCarId <= 0 {
		return types.ParamLessZero, errors.New("购物车ID 不能小于等于 0")
	}
	if e.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户 ID 不能小于等于 0")
	}
	if e.BuyNums <= 0 {
		return types.ParamLessZero, errors.New("购买数量不能小于等于 0")
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

type DelGoodCarCheck struct {
	GoodsIds []int64 `json:"goods_ids"`
}

func (c DelGoodCarCheck) DelGoodCarCheckParamValidate() (int, error) {
	if c.GoodsIds == nil {
		return types.ParamLessZero, errors.New("商品 ID 数组长度不能小于等于0")
	}
	return types.ReturnSuccess, nil
}

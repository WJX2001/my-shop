package api

import (
	"encoding/json"
	"my-ganji-app/models"
	"my-ganji-app/types"
	type_goods_car "my-ganji-app/types/goods_car"
)

type GoodsCarController struct {
	BaseController
}

// AddGoodsToCar 将商品添加到购物车
func (c *GoodsCarController) AddGoodsToCar() {
	user_token, ok := c.CurrentUser()
	if !ok {
		return
	}

	goods_car := type_goods_car.AddGoodCarCheck{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &goods_car); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式，请联系客服处理")
		c.ServeJSON()
		return
	}
	if code, err := goods_car.AddGoodCarCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, nil, err.Error())
		c.ServeJSON()
		return
	}
	if user_token.Id != goods_car.UserId {
		c.Data["json"] = RetResource(false, types.UserTokenUserIdNotEqual, nil, "传入的用户ID和用户Token不相符")
		c.ServeJSON()
		return
	}
	goods_dtl, _, _ := models.GetGoodsDetail(goods_car.GoodsId)

	// 进行价格校验
	var unitPrice float64
	// 1=非打折，2=打折
	switch goods_car.IsDis {
	case 1:
		unitPrice = goods_dtl.GoodsPrice
	case 2:
		unitPrice = goods_dtl.GoodsDisPrice
	default:
		c.Data["json"] = RetResource(false, types.InvalidVerifyWay, nil, "无效的商品价格方式")
		c.ServeJSON()
		return
	}

	// 统一校验逻辑，只写一次
	expectAmount := unitPrice * float64(goods_car.BuyNums)
	if goods_car.PayAmount != expectAmount {
		c.Data["json"] = RetResource(false, types.InvalidGoodsPrice, nil, "无效的商品价格")
		c.ServeJSON()
		return
	}

	if goods_dtl != nil {
		gdsc, _, _ := models.GetGoodsCarDetailByGoodsId(user_token.Id, goods_dtl.Id)
		if gdsc == nil { // 用户的购物车里面没有这个商品
			gdc := models.GoodsCar{
				GoodsId:    goods_dtl.Id,
				Logo:       goods_dtl.Logo,
				MerchantId: goods_dtl.Id,
				GoodsTypes: goods_car.GoodsTypes,
				GoodsTitle: goods_dtl.Title,
				GoodsName:  goods_dtl.GoodsName,
				UserId:     user_token.Id,
				BuyNums:    goods_car.BuyNums,
				PayAmount:  goods_car.PayAmount,
			}
			err := gdc.Insert()
			if err != nil {
				c.Data["json"] = RetResource(false, types.SystemDbErr, nil, "数据库操作错误")
				c.ServeJSON()
				return
			}
		} else { // 用户的购物车里有这个商品
			gdsc.BuyNums = gdsc.BuyNums + goods_car.BuyNums
			gdsc.PayAmount = gdsc.PayAmount + goods_car.PayAmount
			err := gdsc.Update()
			if err != nil {
				c.Data["json"] = RetResource(false, types.SystemDbErr, nil, "数据库操作错误")
				c.ServeJSON()
				return
			}
		}
		c.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "添加购物车成功")
		c.ServeJSON()
		return
	}
}

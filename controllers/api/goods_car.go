package api

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
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

// GoodsCarList 获取购物车列表
func (c *GoodsCarController) GoodsCarList() {
	ut, ok := c.CurrentUser()
	if !ok {
		return
	}

	goods_car_lst := type_goods_car.GoodCarListCheck{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &goods_car_lst); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式，请联系客服处理")
		c.ServeJSON()
		return
	}
	if code, err := goods_car_lst.GoodCarListCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, nil, err.Error())
		c.ServeJSON()
		return
	}
	goods_car_list, total, err := models.GetGoodsCarList(goods_car_lst.Page, goods_car_lst.PageSize, ut.Id)
	if err != nil {
		c.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, err.Error())
		c.ServeJSON()
		return
	}
	var gds_car_lst []type_goods_car.GoodsCarList
	img_path, _ := beego.AppConfig.String("img_root_path")
	for _, value := range goods_car_list {
		gds, _, _ := models.GetGoodsDetail(value.GoodsId)
		mct, _, _ := models.GetMerchantDetail(gds.MerchantId)
		var goods_price float64
		if gds.IsDiscount == 0 {
			goods_price = gds.GoodsPrice
		} else {
			goods_price = gds.GoodsDisPrice
		}
		gdsc := type_goods_car.GoodsCarList{
			MerchantId:   gds.MerchantId,
			MerchantName: mct.MerchantName,
			GoodsCarId:   value.Id,
			GoodsId:      gds.Id,
			GoodsLogo:    img_path + value.Logo,
			GoodsTitle:   gds.Title,
			GoodsMark:    gds.GoodsMark,
			GoodsName:    gds.GoodsName,
			GoodsPrice:   goods_price,
			UserId:       value.UserId,
			BuyNums:      value.BuyNums,
			PayAmount:    value.PayAmount,
		}
		gds_car_lst = append(gds_car_lst, gdsc)
	}
	data := map[string]interface{}{
		"total":       total,
		"gds_car_lst": gds_car_lst,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取购物车列表成功")
	c.ServeJSON()
	return
}

package api

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"my-ganji-app/models"
	"my-ganji-app/types"
	type_goods "my-ganji-app/types/goods"
)

type GoodsController struct {
	beego.Controller
}

// GoodsCategoryList 分类商品列表接口
func (c *GoodsController) GoodsCategoryList() {
	goods_category := type_goods.GoodsCategoryCheck{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &goods_category); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		c.ServeJSON()
		return
	}
	if code, err := goods_category.GoodsCategoryCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, nil, err.Error())
		c.ServeJSON()
		return
	}
	good_list, total, err := models.GetCategoryGoodsList(goods_category.Page, goods_category.PageSize, goods_category.FirstLevelCatId, goods_category.LastLevelCatId)
	if err != nil {
		c.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, "获取商品列表失败")
		c.ServeJSON()
		return
	}

	image_path, err := beego.AppConfig.String("img_root_path")
	if err != nil {
		c.Data["json"] = RetResource(false, types.InvalidConfig, nil, "解析环境变量错误")
		c.ServeJSON()
		return
	}

	var goods_ret_list []type_goods.CategoryGoodsRet
	for _, value := range good_list {
		gds_ret := type_goods.CategoryGoodsRet{
			GoodsId:       int64(value.Id),
			GoodsMark:     value.GoodsMark,
			Title:         value.Title,
			Logo:          image_path + value.Logo,
			GoodsPrice:    value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			GoodsIntegral: value.GoodsIntegral,
			SendIntegral:  value.SendIntegral,
			LeftTime:      value.LeftTime,
			IsHot:         value.IsHot,
			IsDiscount:    value.IsDiscount,
			IsIgSend:      value.IsIgSend,
			IsGroup:       value.IsGroup,
			IsIntegral:    value.IsIntegral,
		}
		goods_ret_list = append(goods_ret_list, gds_ret)
	}
	data := map[string]interface{}{
		"total":   total,
		"gds_lst": goods_ret_list,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取分类商品列表成功")
	c.ServeJSON()
	return
}

// MerchantGoodsList 商家商品列表接口
func (c *GoodsController) MerchantGoodsList() {
	merchant_gds := type_goods.MerchantGoodsListCheck{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &merchant_gds); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		c.ServeJSON()
		return
	}
	if code, err := merchant_gds.MerchantGoodsListCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, nil, err.Error())
		c.ServeJSON()
		return
	}

	goods_list, total, err := models.GetMerchantGoodsList(merchant_gds.Page, merchant_gds.PageSize, merchant_gds.MerchantId, merchant_gds.QueryWay)
	if err != nil {
		c.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, "获取商品列表失败")
		c.ServeJSON()
		return
	}
	img_path, _ := beego.AppConfig.String("img_root_path")
	var goods_ret_list []type_goods.CategoryGoodsRet
	for _, value := range goods_list {
		gds_ret := type_goods.CategoryGoodsRet{
			GoodsId:       int64(value.Id),
			GoodsMark:     value.GoodsMark,
			Title:         value.Title,
			Logo:          img_path + value.Logo,
			GoodsPrice:    value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			GoodsIntegral: value.GoodsIntegral,
			SendIntegral:  value.SendIntegral,
			LeftTime:      value.LeftTime,
			IsHot:         value.IsHot,
			IsDiscount:    value.IsDiscount,
			IsIgSend:      value.IsIgSend,
			IsGroup:       value.IsGroup,
			IsIntegral:    value.IsIntegral,
		}
		goods_ret_list = append(goods_ret_list, gds_ret)
	}
	data := map[string]interface{}{
		"total":     total,
		"order_lst": goods_ret_list,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取分类商品列表成功")
	c.ServeJSON()
	return
}

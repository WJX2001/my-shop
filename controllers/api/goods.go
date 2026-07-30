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

// GetLimitTimeGoodsList 限时购买产品列表
func (c *GoodsController) GetLimitTimeGoodsList() {
	lt_gds := type_goods.LTGoodsListCheck{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &lt_gds); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		c.ServeJSON()
		return
	}
	goods_list, total, err := models.GetLtGoodsList(lt_gds.Page, lt_gds.PageSize)
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
		"goods_lst": goods_ret_list,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取限时购买商品列表成功")
	c.ServeJSON()
	return
}

// GetHotGoodsList 获取爆款产品列表
func (c *GoodsController) GetHotGoodsList() {
	lt_gds := type_goods.LTGoodsListCheck{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &lt_gds); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		c.ServeJSON()
		return
	}
	goods_list, total, err := models.GetOrderDownHotGoodsList(lt_gds.Page, lt_gds.PageSize)
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
		"goods_lst": goods_ret_list,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取获取爆款产品列表成功")
	c.ServeJSON()
	return
}

func (c *GoodsController) GoodsDetail() {
	goods_detail := type_goods.GoodsDetailCheck{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &goods_detail); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		c.ServeJSON()
		return
	}
	if code, err := goods_detail.GoodsDetailCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, nil, err.Error())
		c.ServeJSON()
		return
	}

	goods_dtl, code, err := models.GetGoodsDetail(goods_detail.GoodsId)
	if err != nil {
		c.Data["json"] = RetResource(false, code, err.Error(), "获取商品列表失败")
		c.ServeJSON()
		return
	}
	img_path, _ := beego.AppConfig.String("img_root_path")
	merchant, code, err := models.GetMerchantDetail(goods_dtl.MerchantId)
	if err != nil {
		c.Data["json"] = RetResource(false, code, err.Error(), "获取商家信息失败")
		c.ServeJSON()
		return
	}
	merchant_info := map[string]interface{}{
		"merchant_id":   merchant.Id,
		"merchant_logo": img_path + merchant.Logo,
		"merchant_name": merchant.MerchantName,
	}
	goods_img_lst, code, err := models.GetGoodsImgList(goods_dtl.Id)
	if err != nil {
		c.Data["json"] = RetResource(false, code, err.Error(), "获取商品图片失败")
		c.ServeJSON()
		return
	}
	gds_img_lst := []type_goods.GoodsImagesRet{}
	for _, v := range goods_img_lst {
		gds_img := type_goods.GoodsImagesRet{
			GoodsImgId: v.Id,
			ImageUrl:   img_path + v.Image,
		}
		gds_img_lst = append(gds_img_lst, gds_img)
	}
	user_address := make(map[string]interface{})
	if goods_detail.UserId > 0 {
		user_addr, _, err := models.GetUserAddressDefault(goods_detail.UserId)
		if err != nil {
			user_address = nil
		} else {
			user_address["address_id"] = user_addr.Id
			user_address["address_name"] = user_addr.Address
		}
	} else {
		user_address = nil
	}
	type_list_data, _, err := models.GetGoodsTypeList(goods_dtl.Id)
	var type_list []type_goods.GoodsTypeRet
	// 颜色属性
	if err != nil || type_list_data == nil {
		type_list = nil
	} else {
		for _, value_t := range type_list_data {
			var value_list []string
			json.Unmarshal([]byte(value_t.TypeVale), &value_list)
			c_gds_type := type_goods.GoodsTypeRet{
				GdsTypeKey:   value_t.TypeKey,
				GdsTypeValue: value_list,
			}
			type_list = append(type_list, c_gds_type)

		}
	}
	res := map[string]interface{}{
		"id":              goods_dtl.Id,
		"title":           goods_dtl.Title,
		"mark":            goods_dtl.GoodsMark,
		"logo":            img_path + goods_dtl.Logo,
		"serveice":        goods_dtl.Serveice,
		"calc_way":        goods_dtl.CalcWay,
		"sell_nums":       goods_dtl.SellNums,
		"total_amount":    goods_dtl.TotalAmount,
		"left_amount":     goods_dtl.LeftAmount,
		"goods_price":     goods_dtl.GoodsPrice,
		"goods_dis_price": goods_dtl.GoodsDisPrice,
		"goods_integral":  goods_dtl.GoodsIntegral,
		"group_number":    goods_dtl.GroupNumber,
		"send_interal":    goods_dtl.SendIntegral,
		"goods_name":      goods_dtl.GoodsName,
		"goods_params":    goods_dtl.GoodsParams,
		"goods_detail":    goods_dtl.GoodsDetail,
		"goods_img":       gds_img_lst,
		"user_address":    user_address,
		"merchant_info":   merchant_info,
		"is_hot":          goods_dtl.IsHot,
		"is_discount":     goods_dtl.IsDiscount,
		"is_ig_send":      goods_dtl.IsIgSend,
		"is_group":        goods_dtl.IsGroup,
		"is_integral":     goods_dtl.IsIntegral,
		"goods_types":     type_list,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, res, "获取商品详情成功")
	c.ServeJSON()
	return
}

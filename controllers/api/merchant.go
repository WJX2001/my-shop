package api

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"my-ganji-app/models"
	"my-ganji-app/types"
	type_merchant "my-ganji-app/types/merchant"
)

type MerchantController struct {
	beego.Controller
}

// MerchantList 商家列表接口
func (c *MerchantController) MerchantList() error {
	gds_merchant := type_merchant.MerchantListCheck{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &gds_merchant); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		c.ServeJSON()
		return nil
	}
	if code, err := gds_merchant.MerchantListCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, nil, err.Error())
		c.ServeJSON()
		return nil
	}
	merchant_list, total, err := models.GetMerchantList(gds_merchant.Page, gds_merchant.PageSize, gds_merchant.MerchantName, gds_merchant.MerchantAddress)
	if err != nil {
		c.Data["json"] = RetResource(false, types.GetMerchantListFail, nil, "获取商家列表失败")
		c.ServeJSON()
		return nil
	}
	image_path, err := beego.AppConfig.String("img_root_path")
	if err != nil {
		return err
	}
	var mct_list_ret []type_merchant.MerchantListRet
	for _, merchant := range merchant_list {
		mct_ret := type_merchant.MerchantListRet{
			MctId:        merchant.Id,
			MctName:      merchant.MerchantName,
			MctIntroduce: merchant.MerchantIntro,
			MctLogo:      image_path + merchant.Logo,
			MctWay:       merchant.MerchantWay,
			ShopLevel:    merchant.ShopLevel,
			ShopServer:   merchant.ShopServer,
		}
		mct_list_ret = append(mct_list_ret, mct_ret)
	}
	data := map[string]interface{}{
		"total":   total,
		"mct_lst": mct_list_ret,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取商家列表成功")
	c.ServeJSON()
	return nil
}

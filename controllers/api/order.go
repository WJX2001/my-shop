package api

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/google/uuid"
	"my-ganji-app/models"
	"my-ganji-app/types"
	type_order "my-ganji-app/types/order"
)

type OrderController struct {
	BaseController
}

func (c *OrderController) CreateOrder() {
	requestUser, ok := c.CurrentUser()
	if !ok {
		return
	}
	if requestUser == nil {
		c.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		c.ServeJSON()
		return
	}

	var create_order type_order.CreateOrderCheck
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &create_order); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式，请联系客服处理")
		c.ServeJSON()
		return
	}
	if code, err := create_order.CreateOrderCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, nil, err.Error())
		c.ServeJSON()
		return
	}
	if requestUser.Id != create_order.UserId {
		c.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "Token 和用户不匹配")
		c.ServeJSON()
		return
	}

	gds, code, err := models.GetGoodsDetail(create_order.GoodsId)
	if err != nil || gds == nil {
		c.Data["json"] = RetResource(false, code, nil, "商品不存在")
		c.ServeJSON()
		return
	}

	var (
		payAmount    float64
		payIntegral  float64
		sendIntegral float64
	)
	buyNums := float64(create_order.BuyNums)

	// 价格 / 积分 / 赠送积分全部按商品表计算，不信任前端
	if gds.IsIntegral == 1 {
		if create_order.PayWay != 0 {
			c.Data["json"] = RetResource(false, types.InvalidVerifyWay, nil, "积分商品请使用积分支付")
			c.ServeJSON()
			return
		}
		payIntegral = gds.GoodsIntegral * buyNums
		i_gl, err := models.GetIntegralByUserId(requestUser.Id)
		if err != nil || i_gl == nil {
			c.Data["json"] = RetResource(false, types.SystemDbErr, nil, "查询用户积分失败")
			c.ServeJSON()
			return
		}
		if i_gl.TodayIg < payIntegral {
			c.Data["json"] = RetResource(false, types.IntegralNotEnough, nil, "您的账户积分不足")
			c.ServeJSON()
			return
		}
	} else {
		if create_order.PayWay == 0 {
			c.Data["json"] = RetResource(false, types.InvalidVerifyWay, nil, "非积分商品不能使用积分支付")
			c.ServeJSON()
			return
		}
		unitPrice := gds.GoodsPrice
		if gds.IsDiscount == 1 {
			unitPrice = gds.GoodsDisPrice
		}
		payAmount = unitPrice * buyNums
		sendIntegral = gds.SendIntegral * buyNums
	}

	order_nmb := uuid.New().String()
	cmt := models.GoodsOrder{
		GoodsId:       gds.Id,
		MerchantId:    gds.MerchantId,
		AddressId:     create_order.AddressId,
		GoodsTypes:    create_order.GoodsTypes,
		GoodsTitle:    gds.Title,
		GoodsName:     gds.GoodsName,
		Logo:          gds.Logo,
		UserId:        create_order.UserId,
		BuyNums:       create_order.BuyNums,
		PayWay:        create_order.PayWay,
		PayAmount:     payAmount,
		PayIntegral:   payIntegral,
		SendIntegral:  sendIntegral,
		OrderNumber:   order_nmb,
		OrderStatus:   0,
		FailureReason: "未支付",
		BatchId:       order_nmb,
	}
	err, id := cmt.Insert()
	if err != nil {
		c.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "创建订单失败")
		c.ServeJSON()
		return
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, map[string]interface{}{"id": id}, "创建订单成功")
	c.ServeJSON()
}

// OrderList 订单列表 OrderList
func (c *OrderController) OrderList() {
	u_tk, ok := c.CurrentUser()
	if !ok {
		return
	}
	var order_lst type_order.OrderListCheck
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &order_lst); err != nil {
		c.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		c.ServeJSON()
		return
	}
	if code, err := order_lst.OrderListCheckParamValidate(); err != nil {
		c.Data["json"] = RetResource(false, code, err, err.Error())
		c.ServeJSON()
		return
	}
	ols, total, err := models.GetGoodsOrderList(order_lst.Page, order_lst.PageSize, u_tk.Id, order_lst.OrderStatus)
	if err != nil {
		c.Data["json"] = RetResource(false, types.SystemDbErr, err, err.Error())
		c.ServeJSON()
		return
	}
	var olst_ret []type_order.OrderListRet
	img_path, _ := beego.AppConfig.String("img_root_path")
	for _, value := range ols {
		m, _, _ := models.GetMerchantDetail(value.MerchantId)
		gds, _, _ := models.GetGoodsDetail(value.GoodsId)
		var goods_last_price float64
		if gds.IsDiscount == 0 {
			goods_last_price = gds.GoodsPrice
		} else {
			goods_last_price = gds.GoodsDisPrice
		}
		ordr := type_order.OrderListRet{
			MerchantId:    m.Id,
			MerchantName:  m.MerchantName,
			MerchantPhone: m.Phone,
			OrderId:       value.Id,
			GoodsName:     value.GoodsName,
			GoodsLogo:     img_path + value.Logo,
			GoodsPrice:    goods_last_price,
			PayIntegral:   value.PayIntegral,
			SendIntegral:  gds.SendIntegral,
			OrderStatus:   value.OrderStatus,
			BuyNums:       value.BuyNums,
			PayAmount:     value.PayAmount,
			IsCancle:      value.IsCancle,
			IsComment:     value.IsComment,
			IsDiscount:    gds.IsDiscount,
			IsIntegral:    gds.IsIntegral,
		}
		olst_ret = append(olst_ret, ordr)
	}
	data := map[string]interface{}{
		"total":     total,
		"order_lst": olst_ret,
	}
	c.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取订单列表成功")
	c.ServeJSON()
	return
}

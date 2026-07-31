package order

type OrderListRet struct {
	MerchantId    int64   `json:"merchant_id"`
	MerchantName  string  `json:"merchant_name"`
	MerchantPhone string  `json:"merchant_phone"`
	OrderId       int64   `json:"order_id"`
	GoodsName     string  `json:"goods_name"`
	GoodsLogo     string  `json:"goods_logo"`
	GoodsPrice    float64 `json:"goods_price"`
	PayIntegral   float64 `json:"pay_integral"`
	SendIntegral  float64 `json:"send_integral"`
	OrderStatus   int8    `json:"order_status"`
	BuyNums       int64   `json:"buy_nums"`
	PayAmount     float64 `json:"pay_amount"`
	IsCancle      int8    `json:"is_cancle"`
	IsComment     int8    `json:"is_comment"`
	IsDiscount    int8    `json:"is_discount"` // 0:不打折，1:打折活动产品
	IsIntegral    int8    `json:"is_integral"`
}

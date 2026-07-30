package goods

import (
	"github.com/pkg/errors"
	"my-ganji-app/types"
)

type GoodsCategoryCheck struct {
	types.PageSizeData
	FirstLevelCatId int64 `json:"first_levet_cat_id"`
	LastLevelCatId  int64 `json:"last_level_cat_id"`
}

func (gcc GoodsCategoryCheck) GoodsCategoryCheckParamValidate() (int, error) {
	code, err := gcc.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if gcc.FirstLevelCatId < 0 {
		return types.ParamLessZero, errors.New("商品一级分类ID小于0")
	}
	if gcc.LastLevelCatId < 0 {
		return types.ParamLessZero, errors.New("商品二级分类ID小于0")
	}
	return types.ReturnSuccess, nil
}

type MerchantGoodsListCheck struct {
	types.PageSizeData
	MerchantId int64 `json:"merchant_id"`
	QueryWay   int8  `json:"query_way"` // 0:全部；1:活动优选；2:爆款产品
}

func (mglc MerchantGoodsListCheck) MerchantGoodsListCheckParamValidate() (int, error) {
	code, err := mglc.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if mglc.MerchantId <= 0 {
		return types.ParamLessZero, errors.New("商家ID不能小于等于0")
	}
	if mglc.QueryWay < 0 || mglc.QueryWay > 2 {
		return types.InvalidVerifyWay, errors.New("无效的查询方式")
	}
	return types.ReturnSuccess, nil
}

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

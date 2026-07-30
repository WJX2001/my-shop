package merchant

import (
	"github.com/pkg/errors"
	"my-ganji-app/types"
)

type MerchantListCheck struct {
	types.PageSizeData
	MerchantName    string `json:"merchant_name"`
	MerchantAddress string `json:"merchant_address"`
}

func (m MerchantListCheck) MerchantListCheckParamValidate() (int, error) {
	code, err := m.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	return types.ReturnSuccess, nil
}

type MerchantDetailCheck struct {
	MerchantId int64 `json:"merchant_id"`
}

func (m MerchantDetailCheck) MerchantDetailCheckParamValidate() (int, error) {
	if m.MerchantId <= 0 {
		return types.ParamLessZero, errors.New("MerchantId 不能小于 0")
	}
	return types.ReturnSuccess, nil
}

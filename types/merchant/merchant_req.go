package merchant

import "my-ganji-app/types"

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

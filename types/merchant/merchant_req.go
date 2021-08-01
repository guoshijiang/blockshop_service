package merchant

import (
	"blockshop/types"
	"github.com/pkg/errors"
)

type MerchantListReq struct {
	types.PageSizeData
	MerchantName    string `json:"merchant_name"`
	MerchantAddress string `json:"merchant_address"`
}

func (this MerchantListReq) ParamCheck() (int, error) {
	code, err := this.ParamCheck()
	if err != nil {
		return code, err
	}
	return types.ReturnSuccess, nil
}

type MerchantDetailReq struct {
	MerchantId int64 `json:"merchant_id"`
}

func (this MerchantDetailReq) ParamCheck() (int, error) {
	if this.MerchantId <= 0 {
		return types.ParamEmptyError, errors.New("MerchantId 不能小于 0")
	}
	return types.ReturnSuccess, nil
}


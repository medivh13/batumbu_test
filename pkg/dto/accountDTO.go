package dto

import (
	"github.com/medivh13/batumbu_test/pkg/errors"
)

type AccountDTO struct {
	AccountNo    string  `json:"account_number"`
	CustomerName string  `json:"customer_name"`
	Balance      float64 `json:"balance"`
}

type TransferReqDTO struct {
	RequestID string  `json:"request_id"`
	ToAccount string  `json:"to_account_number"`
	Amount    float64 `json:"amount"`
}

func (d *TransferReqDTO) Validate() error {

	if d.RequestID == "" || d.ToAccount == "" || d.Amount == 0 {
		return errors.ErrInvalidRequest
	}
	return nil
}

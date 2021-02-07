package dto

import (
	"github.com/medivh13/batumbu_test/pkg/errors"
)

type InvestReturnDTO struct {
	ChosenInvestType string  `json:"chosen_invest_type"`
	InvestPeriod     int32   `json:"invest_period"`
	TimeHorizon      int32   `json:"time_horizon"`
	InvestReturn     float64 `json:"invest_return"`
}

type InvestReqDTO struct {
	InvestType   string `json:"invest_type"`
	InvestPeriod int32  `json:"invest_period"`
}

func (d *InvestReqDTO) Validate() error {

	if d.InvestType == "" || d.InvestPeriod == 0 {
		return errors.ErrInvalidRequest
	}
	return nil
}

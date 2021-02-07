package services

import (
	"strings"

	"github.com/medivh13/batumbu_test/pkg/dto"
)

type service struct {
}

func NewService() Services {
	return &service{}
}

func (s *service) GetInvestReturn(req *dto.InvestReqDTO) (interface{}, error) {

	data, err := Calculate(req.InvestType, req.InvestPeriod)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Calculate(investType string, period int32) (*dto.InvestReturnDTO, error) {
	var investReturn dto.InvestReturnDTO

	var yearlyReturn float64
	var timeHorizon float64
	var adjustment float64

	switch strings.ToUpper(investType) {
	case "A":
		yearlyReturn = 6.5
		timeHorizon = 5
		break
	case "B":
		yearlyReturn = 7
		timeHorizon = 3
		break
	case "C":
		yearlyReturn = 12
		timeHorizon = 7
		break
	case "D":
		yearlyReturn = 20
		timeHorizon = 8
		break
	default:
		yearlyReturn = 0
		timeHorizon = 0
		break
	}

	if period > 1 {
		for i := 2; i <= int(period); i++ {
			adjustment += 2
		}

		yearlyReturn += adjustment

		if period >= int32(timeHorizon) {
			yearlyReturn += 12.5
		}
	}

	investReturn.InvestReturn = yearlyReturn
	investReturn.ChosenInvestType = strings.ToUpper(investType)
	investReturn.TimeHorizon = int32(timeHorizon)
	investReturn.InvestPeriod = period

	return &investReturn, nil

}

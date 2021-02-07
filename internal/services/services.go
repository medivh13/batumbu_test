package services

import "github.com/medivh13/batumbu_test/pkg/dto"

type Services interface {
	GetInvestReturn(req *dto.InvestReqDTO) (interface{}, error)
}

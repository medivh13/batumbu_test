package services_test

import (
	"testing"

	"github.com/medivh13/batumbu_test/internal/services"
	"github.com/medivh13/batumbu_test/pkg/dto"
	"github.com/stretchr/testify/assert"
)

func TestGetInvestReturn(t *testing.T) {
	mockData := dto.InvestReqDTO{
		InvestType:   "a",
		InvestPeriod: 5,
	}

	t.Run("success", func(t *testing.T) {
		s := services.NewService()

		a, err := s.GetInvestReturn(&mockData)

		assert.NoError(t, err)
		assert.NotNil(t, a)

	})

}

package http

import (
	"net/http"
	"strconv"

	"github.com/medivh13/batumbu_test/pkg/dto"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/medivh13/batumbu_test/internal/services"
	"github.com/medivh13/batumbu_test/pkg/errors"
)

type HttpHandler struct {
	service services.Services
}

func NewHttpHandler(e *echo.Echo, srv services.Services) {
	handler := &HttpHandler{
		srv,
	}

	e.GET("/invest/:type/:period", handler.GetInvestReturn)
}

func (h *HttpHandler) GetInvestReturn(c echo.Context) error {

	var investReqDTO dto.InvestReqDTO

	investReqDTO.InvestType = c.Param("type")
	i, err := strconv.Atoi(c.Param("period"))

	if err != nil {
		log.Error(err.Error())
		return err
	}
	investReqDTO.InvestPeriod = int32(i)

	if investReqDTO.InvestType == "" || investReqDTO.InvestPeriod == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	data, err := h.service.GetInvestReturn(&investReqDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)

}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case errors.ErrInternalServerError:
		return http.StatusInternalServerError
	case errors.ErrNotFound:
		return http.StatusNotFound
	case errors.ErrConflict:
		return http.StatusConflict
	case errors.ErrInvalidRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

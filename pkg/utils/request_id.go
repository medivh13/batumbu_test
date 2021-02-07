package utils

import (
	"github.com/labstack/echo"
	"github.com/rs/xid"
)

const ContextKeyRequestID string = "requestID"

func AssignRequestID(c echo.Context) {

	reqID := xid.New()

	c.Set(ContextKeyRequestID, reqID.String())

}

func GetRequestID(c echo.Context) string {

	reqID := c.Get(ContextKeyRequestID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}

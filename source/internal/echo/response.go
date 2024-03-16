package echo

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

// ArrayResponse 数组类型响应
type ArrayResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Total int64  `json:"total,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func WrapError(err1 error, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	} else if err1 == nil {
		return err2
	} else if err2 == nil {
		return err1
	}

	return errors.New(err2.Error() + ": " + err1.Error())
}

func response(c echo.Context, code int, errCode int, err error, msg string, data any, total int64) error {
	c.Response().Header().Add("X-Request-Id", fmt.Sprintf("%v", c.Get("requestID")))
	c.Set("code", errCode)
	return WrapError(err, c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Data: data, Total: total}))
}

func errorResponse(c echo.Context, code int, errCode int, err error, msg string) error {
	return response(c, code, errCode, err, msg, nil, 0)
}

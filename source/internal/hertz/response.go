package hertz

import (
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
)

// ArrayResponse 数组类型响应
type ArrayResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Total int    `json:"total,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func response(c *app.RequestContext, code int, errCode int, err error, msg string, data any, total int) {
	c.Header("X-Request-Id", c.GetString("requestID"))
	c.JSON(code, &ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})

	if err != nil {
		_ = c.Error(errors.New(msg + ": " + err.Error()))
	} else {
		_ = c.Error(errors.New(msg))
	}
	c.Set("code", errCode)
	c.Abort()
}

func errorResponse(c *app.RequestContext, code int, errCode int, err error, msg string) {
	response(c, code, errCode, err, msg, nil, 0)
}

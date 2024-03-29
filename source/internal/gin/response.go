package gin

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ArrayResponse 数组类型响应
type ArrayResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Total int64  `json:"total,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func response(c *gin.Context, code int, errCode int, err error, msg string, data any, total int64) {
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

func errorResponse(c *gin.Context, code int, errCode int, err error, msg string) {
	response(c, code, errCode, err, msg, nil, 0)
}

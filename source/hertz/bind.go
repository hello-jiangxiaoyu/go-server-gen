package hertz

import (
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
)

var (
	ErrorSubIsNil       = errors.New("sub is nil")
	ErrorInvalidateType = errors.New("invalidate type")
	ErrorTooManyReq     = errors.New("too many req")
)

func SetReq(c *app.RequestContext, obj ...any) error {
	if len(obj) > 1 {
		return ErrorTooManyReq
	}
	if len(obj) == 0 {
		return nil
	}
	if err := c.BindAndValidate(obj); err != nil {
		return err
	}

	return nil
}

func SetReqWithSub(c *app.RequestContext, obj ...any) (int64, error) {
	if len(obj) > 1 {
		return 0, ErrorTooManyReq
	}

	sub, ok := c.Get("sub")
	if !ok {
		return 0, ErrorSubIsNil
	}

	int64Sub, ok := sub.(int64)
	if !ok {
		return 0, ErrorInvalidateType
	}

	if len(obj) == 0 {
		return int64Sub, nil
	}
	if err := c.BindAndValidate(obj); err != nil {
		return 0, err
	}

	return int64Sub, nil
}

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

type Api struct {
	c     *app.RequestContext
	Sub   int64
	Error error
}

func New(c *app.RequestContext) *Api {
	return &Api{c: c}
}

func (a *Api) SetCtx(ctx *app.RequestContext) *Api {
	if a.c == nil {
		a.c = ctx
	}
	return a
}

func (a *Api) BindJson(obj any) *Api {
	if err := a.c.BindAndValidate(obj); err != nil {
		return a.setError(err)
	}
	return a
}

func (a *Api) BindJsonWithSub(obj any) *Api {
	sub := a.c.GetInt64("sub")
	if sub == 0 {
		return a.setError(errors.New("sub is nil"))
	}
	a.Sub = sub
	if err := a.c.BindAndValidate(obj); err != nil {
		return a.setError(err)
	}

	return a
}

func (a *Api) setError(err error) *Api {
	if a.Error == nil {
		a.Error = err
	}
	return a
}

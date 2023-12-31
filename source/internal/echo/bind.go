package echo

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type Api struct {
	c     echo.Context
	Error error
}

func New(c echo.Context) *Api {
	return &Api{c: c}
}

func (a *Api) SetCtx(ctx echo.Context) *Api {
	if a.c == nil {
		a.c = ctx
	}
	return a
}

func (a *Api) BindJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := a.c.Bind(obj); err != nil {
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

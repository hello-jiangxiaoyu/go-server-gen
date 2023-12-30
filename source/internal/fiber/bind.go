package fiber

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	c     *fiber.Ctx
	Error error
}

func New(c *fiber.Ctx) *Api {
	return &Api{c: c}
}

func (a *Api) SetCtx(ctx *fiber.Ctx) *Api {
	if a.c == nil {
		a.c = ctx
	}
	return a
}

var validate = validator.New()

func (a *Api) BindJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := json.Unmarshal(a.c.Body(), obj); err != nil {
		return a.setError(err)
	}
	if err := validate.Struct(obj); err != nil {
		return a.setError(err)
	}

	return a
}

func (a *Api) BindUri(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := a.c.ParamsParser(obj); err != nil {
		return a.setError(err)
	}
	if err := validate.Struct(obj); err != nil {
		return a.setError(err)
	}

	return a
}

func (a *Api) BindUriAndJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("fiber context is nil"))
	}

	if err := json.Unmarshal(a.c.Body(), obj); err != nil {
		return a.setError(err)
	}
	if err := a.c.ParamsParser(obj); err != nil {
		return a.setError(err)
	}

	if err := validate.Struct(obj); err != nil {
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

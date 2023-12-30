package gin

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Api struct {
	c     *gin.Context
	Error error
}

func New(c *gin.Context) *Api {
	return &Api{c: c}
}

func (a *Api) SetCtx(c *gin.Context) *Api {
	a.c = c
	return a
}

func (a *Api) BindJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("gin context should not be nil"))
	}
	if err := a.c.ShouldBindJSON(obj); err != nil {
		return a.setError(err)
	}
	return a
}

func (a *Api) BindUri(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("gin context should not be nil"))
	}
	if err := setUriValue(a.c, obj); err != nil {
		return a.setError(err)
	}
	return a
}

func (a *Api) BindUriAndJson(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("gin context should not be nil"))
	}
	if err := setUriValue(a.c, obj); err != nil {
		return a.setError(err)
	}
	if err := a.c.ShouldBindJSON(obj); err != nil {
		return a.setError(err)
	}
	return a
}

func (a *Api) BindQuery(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("gin context should not be nil"))
	}
	if err := a.c.BindQuery(obj); err != nil {
		return a.setError(err)
	}
	return a
}

func (a *Api) BindForm(obj any) *Api {
	if a.c == nil {
		return a.setError(errors.New("gin context should not be nil"))
	}
	if err := a.c.ShouldBindWith(obj, binding.Form); err != nil {
		return a.setError(err)
	}
	return a
}

func setUriValue(c *gin.Context, obj any) error {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		fd := t.Field(i)

		if tag := fd.Tag.Get("uri"); tag != "" {
			uri := c.Param(tag)
			if uri == "" {
				if required := fd.Tag.Get("binding"); required != "required" {
					continue
				}
				return errors.New(tag + " uri parameter should not be empty")
			}
			field := reflect.ValueOf(obj).Elem().Field(i)
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				num, err := strconv.ParseInt(uri, 10, 64)
				if err != nil {
					return errors.New("failed to parse " + tag + " parameter as integer")
				}
				field.SetInt(num)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				num, err := strconv.ParseUint(uri, 10, 64)
				if err != nil {
					return errors.New("failed to parse " + tag + " parameter as unsigned integer")
				}
				field.SetUint(num)
			case reflect.Float32, reflect.Float64:
				num, err := strconv.ParseFloat(uri, 64)
				if err != nil {
					return errors.New("failed to parse " + tag + " parameter as float")
				}
				field.SetFloat(num)
			case reflect.String:
				field.SetString(uri)
			default:
				return errors.New("unsupported field type for " + tag + " parameter")
			}
		}
	}
	return nil
}

func (a *Api) setError(err error) *Api {
	if a.Error == nil {
		a.Error = err
	}
	return a
}

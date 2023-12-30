package fiber

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ArrayResponse struct {
	Code  uint   `json:"code"`
	Msg   string `json:"msg"`
	Total int    `json:"total"`
	Data  any    `json:"data"`
}

func response(c *fiber.Ctx, code int, errCode uint, msg string, data any, total int, isArray bool) error {
	c.Locals("code", errCode)
	if isArray {
		return c.Status(code).JSON(&ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})
	}

	return c.Status(code).JSON(&Response{Code: errCode, Msg: msg, Data: data})
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

func WithError(err error, msg string) error {
	if err == nil && msg == "" {
		return nil
	} else if err == nil {
		return errors.New(msg)
	} else if msg == "" {
		return err
	}

	return errors.New(msg + ": " + err.Error())
}

func customResponse(c *fiber.Ctx, code int, errCode uint, err error, msg string, isArray []bool) error {
	isArrayFlag := len(isArray) != 0
	var responseData any
	if isArrayFlag {
		responseData = []struct{}{}
	} else {
		responseData = struct{}{}
	}

	err = WithError(err, msg)
	err = WrapError(err, response(c, code, errCode, msg, responseData, 0, isArrayFlag))

	return err
}

func success(c *fiber.Ctx, data any, total int, isArray bool) error {
	return response(c, fiber.StatusOK, CodeSuccess, MsgSuccess, data, total, isArray)
}

func Success(ctx *fiber.Ctx) error {
	return success(ctx, struct{}{}, 0, false)
}
func SuccessWithData(ctx *fiber.Ctx, data any) error {
	return success(ctx, data, 0, false)
}
func SuccessArray(ctx *fiber.Ctx, total int, data any) error {
	return success(ctx, data, total, true)
}

func DoNothing(ctx *fiber.Ctx, msg string, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusAccepted, CodeSuccess, nil, msg, isArray)
}

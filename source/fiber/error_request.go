package fiber

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ErrorRequest 请求参数错误
func ErrorRequest(ctx *fiber.Ctx, err error, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusBadRequest, CodeRequestPara, err, "invalidate para", isArray)
}

// ErrorRequestWithMsg 请求参数错误
func ErrorRequestWithMsg(ctx *fiber.Ctx, err error, msg string, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusBadRequest, CodeRequestPara, err, msg, isArray)
}

// ErrorForbidden 无权访问
func ErrorForbidden(ctx *fiber.Ctx, msg string, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusForbidden, CodeForbidden, nil, msg, isArray)
}

// ErrorInvalidateToken token 无效
func ErrorInvalidateToken(c *fiber.Ctx, err error, isArray ...bool) error {
	return customResponse(c, http.StatusForbidden, CodeInvalidToken, err, "invalidate token", isArray)
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(ctx *fiber.Ctx, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusUnauthorized, CodeNotLogin, nil, "user not login", isArray)
}

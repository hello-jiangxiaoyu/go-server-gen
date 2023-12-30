package fiber

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorUnknown 未知错误
func ErrorUnknown(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusInternalServerError, CodeUnknown, err, respMsg, isArray)
}

// ErrorNotFound 资源未找到
func ErrorNotFound(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusInternalServerError, CodeNotFound, err, respMsg, isArray)
}

func ErrorSaveSession(ctx *fiber.Ctx, err error, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusInternalServerError, CodeSaveSession, err, "failed to save session", isArray)
}

// ErrorSendRequest 发送 fast http 请求失败
func ErrorSendRequest(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusInternalServerError, CodeSendRequest, err, respMsg, isArray)
}

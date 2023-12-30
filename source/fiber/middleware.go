package fiber

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorPanic(ctx *fiber.Ctx) error {
	return customResponse(ctx, fiber.StatusInternalServerError, CodeServerPanic, nil, "server panic", nil)
}

func ErrorHost(ctx *fiber.Ctx) error {
	return customResponse(ctx, fiber.StatusForbidden, CodeNoSuchHost, nil, "no such host", nil)
}

// ErrorNoSuchRoute 404
func ErrorNoSuchRoute(ctx *fiber.Ctx, err error, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusNotFound, CodeNoRoute, err, "no such api route", isArray)
}

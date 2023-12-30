package fiber

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ErrorSqlUpdate SQL更新失败
func ErrorSqlUpdate(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	if err != nil && strings.HasPrefix(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		return customResponse(ctx, fiber.StatusConflict, CodeSqlModifyDuplicate, err, respMsg, isArray)
	}
	return customResponse(ctx, fiber.StatusInternalServerError, CodeSqlModify, err, respMsg, isArray)
}

// ErrorSqlCreate SQL创建失败
func ErrorSqlCreate(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	if err != nil && strings.HasPrefix(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		return customResponse(ctx, fiber.StatusConflict, CodeSqlCreateDuplicate, err, respMsg, isArray)
	}
	return customResponse(ctx, fiber.StatusInternalServerError, CodeSqlCreate, err, respMsg, isArray)
}

// ErrorSelect 数据库查询错误
func ErrorSelect(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	if err == gorm.ErrRecordNotFound { // gorm find操作record not found
		return customResponse(ctx, fiber.StatusNotFound, CodeSqlSelectNotFound, err, respMsg, isArray)
	}
	return customResponse(ctx, fiber.StatusInternalServerError, CodeSqlSelect, err, respMsg, isArray)
}

// ErrorSqlDelete SQL删除失败
func ErrorSqlDelete(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	if err == gorm.ErrForeignKeyViolated { // 外键依赖导致无法删除
		return customResponse(ctx, fiber.StatusConflict, CodeSqlDeleteForKey, err, respMsg, isArray)
	}
	return customResponse(ctx, fiber.StatusInternalServerError, CodeSqlDelete, err, respMsg, isArray)
}

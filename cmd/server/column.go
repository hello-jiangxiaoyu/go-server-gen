package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// GetTableColumns 获取表的列信息
func GetTableColumns(c *gin.Context) {
	columns := make([]Column, 0)
	if err := DB.Raw(`desc ` + c.Param("table")).Scan(&columns).Error; err != nil {
		sendErrorResponse(c, err)
		return
	}
	viewColumn := make([]ViewColumn, 0)
	for _, v := range columns {
		item := ViewColumn{
			Column:      v.Field,
			Label:       v.Field,
			LabelWidth:  100,
			Type:        v.Type,
			Key:         v.Key,
			Placeholder: v.Field,
			Required:    false,
			ViewType:    "string",
		}

		if v.Field == "description" {
			item.Label = "描述"
			item.ViewType = "text"
		} else if v.Field == "avatar" {
			item.Label = "头像"
		} else if v.Field == "gender" {
			item.Label = "性别"
		} else if v.Field == "name" {
			item.Label = "名称"
		} else if v.Field == "phone" {
			item.Label = "手机号"
		}

		if v.Field == "avatar" || strings.HasSuffix(v.Field, "image") || strings.HasSuffix(v.Field, "picture") {
			item.ViewType = "image"
		} else if v.Type == "tinyint(1)" {
			item.ViewType = "switch"
		} else if v.Type == "int" || v.Type == "bigint" {
			item.ViewType = "number"
		} else if v.Type == "timestamp" {
			item.ViewType = "datetime"
		} else if v.Type == "date" {
			item.ViewType = "date"
		} else if v.Type == "json" {
			item.ViewType = "json"
		} else if v.Type == "char(1)" {
			item.ViewType = "select"
		}

		if v.Key == "PRI" {
			item.CanSearch = true
		} else if v.Field != "created_at" && v.Field != "updated_at" && v.Field != "updated_by" && v.Field != "deleted_at" {
			item.CanCreate = true
			item.CanEdit = true
		}

		switch v.Field {
		case "created_at":
			item.Label = "创建时间"
		case "updated_at":
			item.Label = "更新时间"
		case "updated_by":
			item.Label = "操作人"
		case "deleted_at":
			item.Label = "删除时间"
		}

		viewColumn = append(viewColumn, item)
	}
	c.JSON(http.StatusOK, viewColumn)
}

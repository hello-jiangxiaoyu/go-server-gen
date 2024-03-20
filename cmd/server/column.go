package server

import (
	"github.com/gin-gonic/gin"
	"go-server-gen/parse"
	"net/http"
	"strings"
)

// GetTableColumns 获取表的列信息
func GetTableColumns(c *gin.Context) {
	columns := make([]Column, 0)
	if err := DB.Raw(`desc ` + c.Param("table")).Scan(&columns).Error; err != nil {
		SendErrorResponse(c, err)
		return
	}
	viewColumn := make([]parse.ViewColumn, 0)
	for _, v := range columns {
		item := parse.ViewColumn{
			Column:     v.Field,
			Label:      getLabelByField(v.Field),
			LabelWidth: 0,
			Type:       v.Type,
			Key:        v.Key,
			Required:   false,
			ViewType:   getViewTypeByDbType(v.Type),
		}

		if v.Field == "description" {
			item.ViewType = "text"
		} else if v.Field == "avatar" || strings.HasSuffix(v.Field, "image") || strings.HasSuffix(v.Field, "picture") {
			item.ViewType = "image"
		}

		if v.Key == "PRI" {
			item.CanSearch = true
		} else if v.Field != "created_at" && v.Field != "updated_at" && v.Field != "updated_by" && v.Field != "deleted_at" {
			item.CanCreate = true
			item.CanEdit = true
		}

		viewColumn = append(viewColumn, item)
	}
	c.JSON(http.StatusOK, viewColumn)
}

func getLabelByField(columnName string) string {
	dic := map[string]string{
		"name":        "名称",
		"username":    "账号",
		"password":    "密码",
		"description": "描述",
		"avatar":      "头像",
		"gender":      "性别",
		"phone":       "手机号",
		"addr":        "地址",
		"address":     "地址",
		"role_id":     "角色id",
		"user_id":     "用户id",
		"created_at":  "创建时间",
		"updated_at":  "更新时间",
		"updated_by":  "操作人",
		"deleted_at":  "删除时间",
	}
	res, ok := dic[columnName]
	if !ok {
		return columnName
	}
	return res
}

func getViewTypeByDbType(columnName string) string {
	dic := map[string]string{
		"tinyint(1)": "switch",
		"int":        "number",
		"bigint":     "number",
		"timestamp":  "datetime",
		"date":       "date",
		"json":       "json",
		"char(1)":    "select",
	}

	res, ok := dic[columnName]
	if !ok {
		return "string"
	}
	return res
}

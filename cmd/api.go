package cmd

import (
	"github.com/gin-gonic/gin"
	"go-server-gen/conf"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

var db *gorm.DB

type (
	Column struct {
		Field string `json:"Field"`
		Type  string `json:"Type"`
		Key   string `json:"Key"`
	}
	ViewColumn struct {
		Field         string   `json:"field"`
		Type          string   `json:"type"`
		Key           string   `json:"key"`
		ViewType      string   `json:"viewType"` // 前端展示类型
		CanCreate     bool     `json:"canCreate"`
		CanEdit       bool     `json:"canEdit"`
		CanSearch     bool     `json:"canSearch"`
		Required      bool     `json:"required"`
		Placeholder   string   `json:"placeholder"`
		SelectOptions []string `json:"selectOptions"`
	}
)

func getTableColumns(c *gin.Context) {
	columns := make([]Column, 0)
	if err := db.Raw(`desc ` + c.Param("table")).Scan(&columns).Error; err != nil {
		sendErrorResponse(c, err)
		return
	}
	viewColumn := make([]ViewColumn, 0)
	for _, v := range columns {
		item := ViewColumn{
			Field:       v.Field,
			Type:        v.Type,
			Key:         v.Key,
			Placeholder: v.Field,
			Required:    false,
			ViewType:    "string",
		}

		if v.Type == "int" || v.Type == "bigint" {
			item.ViewType = "number"
		} else if v.Type == "timestamp" {
			item.ViewType = "datetime"
		} else if v.Type == "date" {
			item.ViewType = "date"
		} else if v.Type == "json" {
			item.ViewType = "json"
		} else if v.Field == "description" {
			item.ViewType = "text"
		} else if v.Field == "avatar" || strings.HasSuffix(v.Field, "image") || strings.HasSuffix(v.Field, "picture") {
			item.ViewType = "image"
		} else if strings.HasPrefix(v.Field, "is_") || v.Type == "tinyint(1)" {
			item.ViewType = "switch"
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

func registerApi(r *gin.Engine) {
	// 获取mysql table 列表
	r.GET("/api/tables", func(c *gin.Context) {
		result := make([]string, 0)
		if err := db.Raw("SHOW TABLES;").Scan(&result).Error; err != nil {
			sendErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, result)
	})

	// 获取mysql列数据
	r.GET("/api/tables/:table/columns", getTableColumns)
	r.GET("/api/layout", func(c *gin.Context) {
		layout, idl, err := conf.GetConfig(ServerType, LogType, LayoutPath, IdlPath)
		if err != nil {
			sendErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, map[string]any{"layout": layout, "idl": idl})
	})

	// 生成代码
	r.POST("/api/tables/:table/generate", func(c *gin.Context) {
		viewColumn := make([]ViewColumn, 0)
		if err := c.ShouldBindJSON(&viewColumn); err != nil {
			sendErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, struct{}{})
	})

}

func sendErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{"msg": err.Error()})
}

package cmd

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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

func registerApi(r *gin.Engine) {
	// 获取mysql列数据
	r.GET("/api/tables/:table/columns", func(c *gin.Context) {
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
				Required:    true,
			}
			viewColumn = append(viewColumn, item)
		}
		c.JSON(http.StatusOK, viewColumn)
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

	// 获取mysql table 列表
	r.GET("/api/tables", func(c *gin.Context) {
		result := make([]string, 0)
		if err := db.Raw("SHOW TABLES;").Scan(&result).Error; err != nil {
			sendErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, result)
	})
}

func sendErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{"msg": err.Error()})
}

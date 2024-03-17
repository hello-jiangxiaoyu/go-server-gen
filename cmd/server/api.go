package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

type (
	Column struct {
		Field string `json:"Field"`
		Type  string `json:"Type"`
		Key   string `json:"Key"`
	}
	ViewColumn struct {
		Key           string   `json:"key"`
		Column        string   `json:"column"`
		Label         string   `json:"label"`
		LabelWidth    int      `json:"labelWidth"`
		Type          string   `json:"type"`
		ViewType      string   `json:"viewType"` // 前端展示类型
		CanCreate     bool     `json:"canCreate"`
		CanEdit       bool     `json:"canEdit"`
		CanSearch     bool     `json:"canSearch"`
		Required      bool     `json:"required"`
		Placeholder   string   `json:"placeholder"`
		SelectOptions []string `json:"selectOptions"`
	}
)

func sendErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{"msg": err.Error()})
}

// GenCode 开始生成代码
func GenCode(c *gin.Context) {
	viewColumn := make([]ViewColumn, 0)
	if err := c.ShouldBindJSON(&viewColumn); err != nil {
		sendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, struct{}{})
}

func GetTableList(c *gin.Context) {
	result := make([]string, 0)
	if err := DB.Raw("SHOW TABLES;").Scan(&result).Error; err != nil {
		sendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

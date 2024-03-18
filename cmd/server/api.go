package server

import (
	"github.com/gin-gonic/gin"
	"go-server-gen/parse"
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
)

func SendErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{"msg": err.Error()})
}

// GenCode 开始生成代码
func GenCode(c *gin.Context) {
	viewColumn := make([]parse.ViewColumn, 0)
	if err := c.ShouldBindJSON(&viewColumn); err != nil {
		SendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, struct{}{})
}

func GetTableList(c *gin.Context) {
	result := make([]string, 0)
	if err := DB.Raw("SHOW TABLES;").Scan(&result).Error; err != nil {
		SendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

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
)

func SendErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{"msg": err.Error()})
}

func GetTableList(routerPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := make([]string, 0)
		if err := DB.Raw("SHOW TABLES;").Scan(&result).Error; err != nil {
			SendErrorResponse(c, err)
			return
		}
		c.JSON(http.StatusOK, map[string]any{
			"tables":       result,
			"routerPrefix": routerPrefix,
		})
	}
}

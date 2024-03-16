package cmd

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func registerApi(r *gin.Engine) {
	// 获取mysql table 列表
	r.GET("/api/tables", func(c *gin.Context) {
		result := make([]string, 0)
		if err := db.Raw("SHOW TABLES;").Scan(&result).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, result)
	})

	// 获取mysql表结构
	r.GET("/api/tables/:table/columns", func(c *gin.Context) {
		columns := make([]Column, 0)
		if err := db.Raw(`desc ` + c.Param("table")).Scan(&columns).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, columns)
	})

	// 生成代码
	r.POST("/api/tables/:table/generate", func(c *gin.Context) {

	})
}

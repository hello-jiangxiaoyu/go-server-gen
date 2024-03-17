package cmd

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-server-gen/cmd/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"strings"
)

func StartWebServer(_ *cobra.Command, _ []string) {
	if len(GormDSN) == 0 {
		println("dsn is required")
		os.Exit(1)
	}
	var err error
	gin.SetMode(gin.ReleaseMode)
	server.DB, err = gorm.Open(mysql.Open(GormDSN), &gorm.Config{})
	if err != nil {
		println("open db err: ", err.Error())
		os.Exit(1)
	}

	r := gin.Default()
	r.GET("/api/tables", server.GetTableList)
	r.GET("/api/tables/:table/columns", server.GetTableColumns)
	r.POST("/api/tables/:table/generate", server.GenCode)
	r.Use(StaticWebFile()) // 其他静态资源

	r.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	println("\n\nLocal:	http://localhost:8080\n")

	if err = r.Run(":8080"); err != nil {
		os.Exit(1)
	}
}

//go:embed web
var code embed.FS

// StaticWebFile web网页代理
func StaticWebFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := "web" + c.Request.URL.Path
		if p == "web/" {
			p = "web/index.html"
		}
		res, err := code.Open(p)
		if err != nil {
			return
		}

		if strings.HasSuffix(p, ".html") {
			c.Header("Content-Type", "text/html; charset=utf-8")
		} else if strings.HasSuffix(p, ".css") {
			c.Header("Content-Type", "text/css; charset=utf-8")
		} else if strings.HasSuffix(p, ".js") {
			c.Header("Content-Type", "text/javascript; charset=utf-8")
		}
		c.Status(http.StatusOK)
		c.Stream(func(w io.Writer) bool {
			io.Copy(w, res)
			return false
		})
		c.Abort()
	}
}

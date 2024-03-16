package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path"
	"strings"
)

func StartWebServer(_ *cobra.Command, _ []string) {
	var err error
	gin.SetMode(gin.ReleaseMode)
	db, err = gorm.Open(mysql.Open("gorm:gorm@tcp(127.0.0.1:3306)/uni_test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		println("open db err: ", err.Error())
		os.Exit(1)
	}

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]any{"msg": "no such router"})
	})
	registerApi(r)
	r.Use(StaticWebFile("cmd/web")) // 其他静态资源
	if err = r.Run(":8080"); err != nil {
		os.Exit(1)
	}
}

// StaticWebFile web网页代理
func StaticWebFile(root string) gin.HandlerFunc {
	type localFileSystem struct {
		http.FileSystem
		root    string
		indexes bool
	}
	fs := &localFileSystem{
		FileSystem: gin.Dir(root, false),
		root:       root,
		indexes:    false,
	}
	fileServer := http.FileServer(fs)

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			return
		}

		if p := strings.TrimPrefix(c.Request.URL.Path, "/"); len(p) < len(c.Request.URL.Path) {
			name := path.Join(fs.root, p)
			stats, err := os.Stat(name)
			if err != nil {
				return
			}
			if stats.IsDir() {
				if !fs.indexes {
					index := path.Join(name, "index.html")
					if _, err = os.Stat(index); err != nil {
						return
					}
				}
			}
			req := c.Request.Clone(c)
			fileServer.ServeHTTP(c.Writer, req)
			c.Abort()
		}
	}
}

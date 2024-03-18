package server

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

// StaticWebFile web网页代理
func StaticWebFile(code *embed.FS) gin.HandlerFunc {
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

// StaticWebFile2 web网页代理
func StaticWebFile2() gin.HandlerFunc {
	root := "cmd/web"
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

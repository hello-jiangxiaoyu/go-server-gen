package cmd

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-server-gen/cmd/server"
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/parse"
	"go-server-gen/writer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func StartWebServer(_ *cobra.Command, _ []string) {
	if len(GormDSN) == 0 {
		println("dsn is required")
		os.Exit(1)
	}
	var err error
	server.DB, err = gorm.Open(mysql.Open(GormDSN), &gorm.Config{})
	if err != nil {
		println("open db err: ", err.Error())
		os.Exit(1)
	}
	autoModifyArgs()
	if err = conf.ReadConfig(LayoutPath, ""); err != nil {
		println("ReadConfig err: ", err.Error())
		os.Exit(1)
	}

	// 启动web服务
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/api/tables", server.GetTableList)
	r.GET("/api/tables/:table/columns", server.GetTableColumns)
	r.POST("/api/tables/:table/generate", GenCode)
	r.Use(server.StaticWebFile(&code)) // 其他静态资源

	println("\n\nLocal:	http://localhost:8080\n")
	if err = r.Run(":8080"); err != nil {
		os.Exit(1)
	}
}

func GenCode(c *gin.Context) {
	req := make([]parse.ViewColumn, 0)
	if err := c.ShouldBindJSON(&req); err != nil {
		server.SendErrorResponse(c, err)
		return
	}

	// 获取配置文件
	idl := parse.GetIdlConfig(c.Param("table"), req)
	layout, _, err := conf.GetConfig(ServerType, LogType, LayoutPath, IdlPath)
	if err != nil {
		server.SendErrorResponse(c, err)
		return
	}

	// 生成数据
	services, messages, err := data.ConfigToData(layout, idl)
	if err != nil {
		server.SendErrorResponse(c, err)
		return
	}

	// 使用数据解析模板
	codeMap := make(map[string]writer.WriteCode)
	if err = parse.GenServiceCode(layout, services, codeMap); err != nil {
		server.SendErrorResponse(c, err)
		return
	}
	if err = parse.GenMessageCode(layout, messages, codeMap); err != nil {
		server.SendErrorResponse(c, err)
		return
	}

	// 将代码写入文件
	if err = writer.Write(codeMap); err != nil {
		server.SendErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, struct{}{})
}

//go:embed web
var code embed.FS

package cmd

import (
	"github.com/spf13/cobra"
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/parse"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"os"
	"strings"
)

func UpdateProject(_ *cobra.Command, args []string) {
	checkUpdateCmdArgs(args)

	// 获取配置文件
	layout, idl, err := conf.GetConfig(ServerType, LogType, LayoutPath, IdlPath)
	if err != nil {
		utils.Log("get config err: ", err.Error())
		os.Exit(1)
	}

	// 生成数据
	services, err := data.ConfigToData(layout, idl)
	if err != nil {
		os.Exit(1)
	}

	// 使用数据解析模板
	codeMap := make(map[string]writer.WriteCode)
	if err = parse.GenServiceCode(layout, services, codeMap); err != nil {
		os.Exit(1)
	}

	// 将代码写入文件
	if err = writer.Write(codeMap); err != nil {
		utils.Log("write error: ", err.Error())
		os.Exit(1)
	}

	println("Success")
}

func checkUpdateCmdArgs(args []string) {
	if len(args) == 0 || args[0] == "" {
		utils.Log("idl path should not be empty")
		os.Exit(1)
	}
	IdlPath = args[0]
	autoModifyArgs()
	if ServerType == "" && LayoutPath == "" {
		utils.Log("server type is empty")
		os.Exit(1)
	}
}

func autoModifyArgs() {
	goMod, err := os.ReadFile("go.mod")
	if err != nil {
		return
	}

	if ServerType == "" {
		if strings.Contains(string(goMod), "github.com/gin-gonic/gin") {
			ServerType = "gin"
		} else if strings.Contains(string(goMod), "github.com/gofiber/fiber") {
			ServerType = "fiber"
		} else if strings.Contains(string(goMod), "github.com/labstack/echo") {
			ServerType = "echo"
		} else if strings.Contains(string(goMod), "github.com/cloudwego/hertz") {
			ServerType = "hertz"
		}
	}
	if OutputDir != "" && !strings.HasSuffix(OutputDir, "/") && !strings.HasSuffix(OutputDir, "\\") {
		OutputDir += "/"
	}
}

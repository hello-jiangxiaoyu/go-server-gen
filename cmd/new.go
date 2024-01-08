package cmd

import (
	"github.com/spf13/cobra"
	"go-server-gen/conf"
	"go-server-gen/source"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"os"
	"os/exec"
	"strings"
)

func NewProject(_ *cobra.Command, args []string) {
	checkNewCmdArgs(args)

	// 新建项目
	if !utils.FileExists(OutputDir + "go.mod") {
		cmd := exec.Command("go", "mod", "init", CreateProjectName)
		if OutputDir != "" {
			if err := os.MkdirAll(OutputDir, os.ModePerm); err != nil {
				utils.Log("make dir err:", err.Error())
				os.Exit(1)
			}
			cmd.Dir = OutputDir
		}
		if _, err := cmd.Output(); err != nil {
			utils.Log("create project err: ", err.Error())
			os.Exit(1)
		}
	} else if !ForceWrite {
		utils.Log("go.mod is already exists! use --force to overwrite")
		os.Exit(1)
	}

	// 获取全局配置
	CreateProjectName, _ = utils.GetProjectName(OutputDir) // 首次获取
	layout, err := conf.GetLayoutConfig(ServerType, LogType, LayoutPath)
	if err != nil {
		utils.Log("get layout config err: ", err.Error())
		os.Exit(1)
	}

	// 生成代码
	res, err := source.GenPackageCode(layout, OutputDir, ForceWrite)
	if err != nil {
		utils.Log("gen default code err: ", err.Error())
		os.Exit(1)
	}

	// 将代码写入文件
	if err = writer.Write(res); err != nil {
		utils.Log("write err: ", err.Error())
		os.Exit(1)
	}

	println("Success")
}

func checkNewCmdArgs(args []string) {
	if ServerType == "" {
		ServerType = "gin"
	}
	if ServerType != "gin" && ServerType != "fiber" &&
		ServerType != "echo" && ServerType != "hertz" {
		utils.Log("server type" + ServerType + " is not valid")
		os.Exit(1)
	}
	if len(args) == 0 || args[0] == "" {
		utils.Log("new project name should not be empty")
		os.Exit(1)
	}
	CreateProjectName = args[0]
	if OutputDir != "" && !strings.HasSuffix(OutputDir, "/") && !strings.HasSuffix(OutputDir, "\\") {
		OutputDir += "/"
	}
}

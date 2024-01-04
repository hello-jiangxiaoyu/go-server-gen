package cmd

import (
	"github.com/spf13/cobra"
	"go-server-gen/conf"
	"go-server-gen/source"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"os"
	"os/exec"
	"path/filepath"
)

func NewProject(_ *cobra.Command, args []string) {
	checkNewCmdArgs(args)
	if CreateProjectName == "" {
		utils.Log("create project name is not valid")
		os.Exit(1)
	}
	cmd := exec.Command("go", "mod", "init", CreateProjectName)
	if _, err := cmd.Output(); err != nil {
		utils.Log("create project err: ", err.Error())
		os.Exit(1)
	}

	// 获取全局配置
	layout, err := conf.GetLayoutConfig(ServerType, LayoutPath)
	if err != nil {
		utils.Log("get layout config err: ", err.Error())
		os.Exit(1)
	}

	// 生成代码
	res := make(map[string]writer.WriteCode)
	if err = source.GenPackageCode(layout, ServerType, LogType, res); err != nil {
		utils.Log("gen default code err: ", err.Error())
		os.Exit(1)
	}

	// 将代码写入文件
	if err = writer.Write(res, ""); err != nil {
		utils.Log("write err: ", err.Error())
		os.Exit(1)
	}

	println("Success")
}

func checkNewCmdArgs(args []string) {
	if ServerType != "gin" && ServerType != "fiber" &&
		ServerType != "echo" && ServerType != "hertz" {
		utils.Log("server type is not valid")
		os.Exit(1)
	}
	if len(args) == 0 {
		utils.Log("new project name is empty")
		os.Exit(1)
	}
	CreateProjectName = args[0]
	if matches, _ := filepath.Glob("go.mod"); len(matches) != 0 {
		utils.Log("go.mod is already exists!")
		os.Exit(1)
	}
}

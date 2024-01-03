package cmd

import (
	"github.com/spf13/cobra"
	"go-server-gen/conf"
	"go-server-gen/source"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"os"
	"os/exec"
)

func NewProject(_ *cobra.Command, _ []string) {
	if ServerType != "gin" && ServerType != "fiber" &&
		ServerType != "echo" && ServerType != "hertz" {
		println("server type is not valid")
		os.Exit(1)
	}

	name, err := utils.GetProjectName()
	if err != nil {
		cmd := exec.Command("go", "mod", "init", CreateProjectName)
		if _, err = cmd.Output(); err != nil {
			println("create project err: ", err.Error())
			os.Exit(1)
		}
	} else {
		CreateProjectName = name // 项目已创建
	}

	// 获取全局配置
	layout, err := conf.GetLayoutConfig(ServerType, LayoutPath)
	if err != nil {
		os.Exit(1)
	}

	// 生成代码
	res := make(map[string]writer.WriteCode)
	if err = source.GenPackageCode(layout, ServerType, LogType, res); err != nil {
		println("gen default code err: ", err.Error())
		os.Exit(1)
	}

	// 将代码写入文件
	if err = writer.Write(res); err != nil {
		println("write err: ", err.Error())
		os.Exit(1)
	}

	println("Success")
}

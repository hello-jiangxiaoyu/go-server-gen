package cmd

import (
	"github.com/spf13/cobra"
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/parse"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"os"
)

func UpdateProject(_ *cobra.Command, _ []string) {
	if IdlPath == "" {
		println("idl path is empty")
		os.Exit(1)
	}

	if err := ExecuteUpdate(ServerType, LayoutPath, IdlPath, ""); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println("Success")
}

func ExecuteUpdate(server, layoutPath, idlPath string, prefix string) error {
	// 获取配置文件
	layout, idl, err := conf.GetConfig(server, idlPath, layoutPath)
	if err != nil {
		utils.Log("get config err: ", err.Error())
		return err
	}

	// 生成数据
	services, messages, err := data.ConfigToData(layout, idl)
	if err != nil {
		return err
	}

	// 使用数据解析模板
	code := make(map[string]writer.WriteCode)
	if err = parse.GenServiceCode(layout, services, code); err != nil {
		return err
	}
	if err = parse.GenMessageCode(layout, messages, code); err != nil {
		return err
	}

	// 将代码写入文件
	if err = writer.Write(code, prefix); err != nil {
		return utils.WithMessage(err, "failed to write code")
	}

	return nil
}

package gen

import (
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/utils"
	"go-server-gen/writer"
)

func ExecuteUpdate(server, layoutPath, idlPath string, prefix string) error {
	// 获取配置文件
	layout, idl, err := conf.GetConfig(server, idlPath, layoutPath)
	if err != nil {
		return utils.WithMessage(err, "failed to unmarshal yaml")
	}

	// 生成数据
	services, messages, err := data.ConfigToData(layout, idl)
	if err != nil {
		return utils.WithMessage(err, "config to data err")
	}

	// 使用数据解析模板
	code := make(map[string]writer.WriteCode)
	if err = ParseServiceCode(layout, services, code); err != nil {
		return err
	}
	if err = ParseMessageCode(layout, messages, code); err != nil {
		return err
	}

	// 将代码写入文件
	if err = writer.Write(code, prefix); err != nil {
		return utils.WithMessage(err, "failed to write code")
	}

	return nil
}

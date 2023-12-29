package gen

import (
	"go-server-gen/conf"
	"go-server-gen/gen/data"
	"go-server-gen/gen/parse"
	"go-server-gen/utils"
	"go-server-gen/writer"
)

func ExecuteUpdate() error {
	// 获取配置文件
	layout, idl, err := conf.GetConfig()
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
	if err = parse.GenServiceCode(layout, services, code); err != nil {
		return err
	}
	if err = parse.GenMessageCode(layout, messages, code); err != nil {
		return err
	}

	// 将代码写入文件
	if err = writer.Write(code); err != nil {
		return utils.WithMessage(err, "failed to write code")
	}

	return nil
}

func ExecuteCreate() error {
	return nil
}

package gen

import (
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/utils"
	"go-server-gen/writer"
)

type GlobalData struct {
	ProjectName   string
	IdlName       string
	HasMiddleware bool
	Pkg           map[string]conf.Package
	Handlers      map[string]string
}

func ParseServiceCode(layout conf.LayoutConfig, services []data.Service, code map[string]writer.WriteCode) error {
	// 全局模板解析
	for _, tpl := range layout.GlobalTemplate {
		handlers := make(map[string]string)
		hasMiddleware := false
		for _, service := range services {
			funcName, handler, err := utils.ParseSource(tpl.HandlerKey, tpl.Handler, service)
			if err != nil {
				return utils.WithMessage(err, "failed to parse and format service handler template")
			}
			handlers[funcName] = handler
			hasMiddleware = hasMiddleware || service.HasMiddleware
		}
		globalData := GlobalData{
			ProjectName:   layout.ProjectName,
			HasMiddleware: hasMiddleware,
			Pkg:           layout.Pkg,
			Handlers:      handlers,
		}
		file, body, err := utils.ParseSource(tpl.Path, tpl.Body, globalData)
		if err != nil {
			return utils.WithMessage(err, "failed to phase and format service body template "+tpl.Path)
		}

		code[file] = writer.WriteCode{
			File:     file,
			Write:    tpl.Write,
			Handlers: handlers,
			Code:     body,
		}
	}

	// service模板解析，一个service对应一个解析文件
	for _, tpl := range layout.ServiceTemplate {
		for _, service := range services {
			handlers := make(map[string]string)
			for _, api := range service.Apis {
				funcName, handler, err := utils.ParseSource(tpl.HandlerKey, tpl.Handler, api)
				if err != nil {
					return utils.WithMessage(err, "failed to parse api handler template")
				}
				handlers[funcName] = handler
			}
			service.Handlers = handlers
			file, body, err := utils.ParseSource(tpl.Path, tpl.Body, service)
			if err != nil {
				return utils.WithMessage(err, "failed to phase and format api body template "+tpl.Path)
			}
			code[file] = writer.WriteCode{
				File:     file,
				Write:    tpl.Write,
				Handlers: handlers,
				Code:     body,
			}
		}
	}

	return nil
}

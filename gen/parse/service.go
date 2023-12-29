package parse

import (
	"go-server-gen/conf"
	"go-server-gen/gen/data"
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

func GenServiceCode(layout *conf.LayoutConfig, services []data.Service) (map[string]writer.WriteCode, error) {
	res := make(map[string]writer.WriteCode)
	// 全局模板解析
	for _, tpl := range layout.GlobalTemplate {
		handlers := make(map[string]string)
		hasMiddleware := false
		for _, service := range services {
			funcName, err := utils.PhaseTemplate(tpl.HandlerKey, service)
			if err != nil {
				return nil, err
			}
			handler, err := utils.PhaseAndFormat(tpl.Handler, service)
			if err != nil {
				return nil, err
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
		body, err := utils.PhaseAndFormat(tpl.Body, globalData)
		if err != nil {
			return nil, utils.WithMessage(err, "failed to phase and format body tpl "+tpl.Name)
		}
		file, err := utils.PhaseTemplate(tpl.Path, globalData)
		if err != nil {
			return nil, utils.WithMessage(err, "failed to phase and format path tpl "+tpl.Path)
		}

		res[file] = writer.WriteCode{
			File:     file,
			Write:    tpl.Write,
			Handlers: handlers,
			Code:     body,
		}
	}

	// service模板解析
	for _, tpl := range layout.ServiceTemplate {
		for _, service := range services {
			handlers := make(map[string]string)
			for _, api := range service.Apis {
				funcName, err := utils.PhaseTemplate(tpl.HandlerKey, api)
				if err != nil {
					return nil, err
				}
				handler, err := utils.PhaseAndFormat(tpl.Handler, api)
				if err != nil {
					return nil, err
				}
				handlers[funcName] = handler
			}
			service.Handlers = handlers
			body, err := utils.PhaseAndFormat(tpl.Body, service)
			if err != nil {
				return nil, utils.WithMessage(err, "failed to phase and format body tpl "+tpl.Name)
			}
			file, err := utils.PhaseTemplate(tpl.Path, service)
			if err != nil {
				return nil, utils.WithMessage(err, "failed to phase and format path tpl "+tpl.Path)
			}
			res[file] = writer.WriteCode{
				File:     file,
				Write:    tpl.Write,
				Handlers: handlers,
				Code:     body,
			}
		}
	}

	return res, nil
}

package parse

import (
	"errors"
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/utils"
	"go-server-gen/writer"
)

func GenServiceCode(layout conf.LayoutConfig, services []data.Service, code map[string]writer.WriteCode) error {
	// 全局模板解析，一个idl对应一个文件
	for _, tpl := range layout.GlobalTemplate {
		writeCode, err := parseGlobal(services, tpl)
		if err != nil {
			return err
		}
		code[writeCode.File] = writeCode
	}

	// service模板解析，一个service对应一个解析文件
	for _, tpl := range layout.ServiceTemplate {
		for _, service := range services {
			writeCode, err := parseService(service, tpl)
			if err != nil {
				return err
			}
			code[writeCode.File] = writeCode
		}
	}

	return nil
}

// 全局模板解析，一个idl对应一个文件
func parseGlobal(services []data.Service, tpl conf.Template) (writer.WriteCode, error) {
	if len(services) == 0 {
		return writer.WriteCode{}, errors.New("service is empty")
	}
	handlers := make(map[string]string)
	for _, service := range services {
		funcName, handler, err := utils.ParseSource(tpl.HandlerKey, tpl.Handler, service)
		if err != nil {
			utils.Log("failed to parse service handler: ", err.Error())
			return writer.WriteCode{}, err
		}
		handlers[funcName] = handler
	}

	file, body, err := utils.ParseSource(tpl.Path, tpl.Body, services[0])
	if err != nil {
		utils.Log("failed to phase "+tpl.Path+" body: ", err.Error())
		return writer.WriteCode{}, err
	}

	return writer.WriteCode{
		File:     file,
		Write:    tpl.Write,
		Handlers: handlers,
		Code:     body,
	}, nil
}

// service模板解析，一个service对应一个解析文件
func parseService(service data.Service, tpl conf.Template) (writer.WriteCode, error) {
	handlers := make(map[string]string)
	for _, api := range service.Apis {
		funcName, handler, err := utils.ParseSource(tpl.HandlerKey, tpl.Handler, api)
		if err != nil {
			utils.Log("failed to parse api handler: ", err.Error())
			return writer.WriteCode{}, err
		}
		handlers[funcName] = handler
	}
	service.Handlers = handlers
	file, body, err := utils.ParseSource(tpl.Path, tpl.Body, service)
	if err != nil {
		utils.Log("failed to parse api body: "+tpl.Path+": ", err.Error())
		return writer.WriteCode{}, err
	}

	return writer.WriteCode{
		File:     file,
		Write:    tpl.Write,
		Handlers: handlers,
		Code:     body,
	}, nil
}

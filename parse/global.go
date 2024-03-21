package parse

import (
	"errors"
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/template"
	"go-server-gen/utils"
	"go-server-gen/writer"
)

// 全局模板解析，一个idl对应一个文件
func parseGlobal(services []data.Service, tpl conf.Template) (writer.WriteCode, error) {
	if len(services) == 0 {
		return writer.WriteCode{}, errors.New("service is empty")
	}
	handlers := make(map[string]string)
	for _, service := range services {
		funcName, handler, err := template.ParseSource(tpl.HandlerKey, tpl.Handler, service)
		if err != nil {
			utils.Logf("failed to parse service handler: %s\n%s\n%s", err.Error(), tpl.HandlerKey, tpl.Handler)
			return writer.WriteCode{}, err
		}
		handlers[funcName] = handler
	}

	globalData := services[0]
	globalData.Handlers = handlers
	file, body, err := template.ParseSource(tpl.Path, tpl.Body, globalData)
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

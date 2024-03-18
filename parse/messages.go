package parse

import (
	"go-server-gen/conf"
	"go-server-gen/data"
	"go-server-gen/utils"
	"go-server-gen/writer"
)

func GenMessageCode(layout conf.LayoutConfig, messages map[string]data.Message, code map[string]writer.WriteCode) error {
	for _, tpl := range layout.MessageTemplate {
		handlers := make(map[string]string)
		for k, msg := range messages {
			_, handler, err := ParseSource("", tpl.Handler, msg)
			if err != nil {
				utils.Log("parse message handler err: ", err.Error())
				return err
			}
			handlers[k] = handler
		}
		file, body, err := ParseSource(tpl.Path, tpl.Body, map[string]any{
			"ProjectName": layout.ProjectName,
			"IdlName":     layout.IdlName,
			"Pkg":         layout.Pkg,
			"Handlers":    handlers,
		})
		if err != nil {
			utils.Log("parse message body err: ", err.Error())
			return err
		}

		code[file] = writer.WriteCode{
			File:     file,
			Write:    tpl.Write,
			Handlers: handlers,
			Code:     body,
		}
	}
	return nil
}

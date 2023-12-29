package parse

import (
	"go-server-gen/conf"
	"go-server-gen/gen/data"
	"go-server-gen/utils"
	"go-server-gen/writer"
)

func GenMessageCode(layout *conf.LayoutConfig, messages map[string]data.Message, code map[string]writer.WriteCode) error {
	for _, tpl := range layout.MessageTemplate {
		handlers := make(map[string]string)
		for k, msg := range messages {
			handler, err := utils.PhaseAndFormat(tpl.Handler, msg)
			if err != nil {
				return err
			}
			handlers[k] = handler
		}
		globalData := GlobalData{
			ProjectName: layout.ProjectName,
			IdlName:     layout.IdlName,
			Pkg:         layout.Pkg,
			Handlers:    handlers,
		}
		body, err := utils.PhaseAndFormat(tpl.Body, globalData)
		if err != nil {
			return utils.WithMessage(err, "failed to phase and format body tpl "+tpl.Name)
		}
		file, err := utils.PhaseTemplate(tpl.Path, globalData)
		if err != nil {
			return utils.WithMessage(err, "failed to phase and format path tpl "+tpl.Path)
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

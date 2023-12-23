package code

import (
	"go-server-gen/gen/conf"
	"go-server-gen/gen/phase"
	"go-server-gen/utils"
)

func GenGroupCode(layout *conf.LayoutConfig, groups []phase.Service) (map[string]WriteCode, error) {
	res := make(map[string]WriteCode)

	for _, tpl := range layout.Templates.Api {
		for _, group := range groups {
			handlers := make(map[string]string)
			if tpl.HandlerTpl != "" {
				funcName, err := utils.PhaseTemplate(tpl.HandlerTpl, group)
				if err != nil {
					return nil, err
				}
				handler, err := utils.PhaseAndFormat(tpl.Handler, group)
				if err != nil {
					return nil, err
				}
				handlers[string(funcName)] = handler
			} else {
				for _, api := range group.Apis {
					handler, err := utils.PhaseAndFormat(tpl.Handler, api)
					if err != nil {
						return nil, utils.WithMessage(err, "failed to phase and format handler tpl "+tpl.Name)
					}
					handlers[api.FuncName] = handler
				}
			}

			group.Handlers = handlers
			body, err := utils.PhaseAndFormat(tpl.Body, group)
			if err != nil {
				return nil, utils.WithMessage(err, "failed to phase and format body tpl "+tpl.Name)
			}
			file, err := utils.PhaseTemplate(tpl.Path, group)
			if err != nil {
				return nil, utils.WithMessage(err, "failed to phase and format path tpl "+tpl.Path)
			}

			res[tpl.Path] = WriteCode{
				File:     string(file),
				Write:    tpl.Write,
				Handlers: handlers,
				Code:     body,
			}
		}
	}

	return res, nil
}

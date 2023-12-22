package code

import (
	"go-server-gen/gen/phase"
	"go-server-gen/utils"
)

func GenGroupCode(g phase.Group) (map[string]string, error) {
	res := make(map[string]string)
	for _, api := range g.Apis {
		handlerCode, err := utils.PhaseAndFormat(g.ApiTemplate, api)
		if err != nil {
			return nil, err
		}
		res[api.FuncName] = handlerCode
		println(handlerCode)
	}

	return res, nil
}

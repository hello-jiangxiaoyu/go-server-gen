package data

import (
	"go-server-gen/conf"
)

// 解析service
func getService(layout conf.LayoutConfig, services []conf.Service, msg map[string]Message) ([]Service, error) {
	res := make([]Service, 0)
	for _, svc := range services {
		group := Service{
			ServiceName:   svc.Name,
			Middleware:    svc.Middlewares,
			Apis:          make([]Api, 0),
			HasMiddleware: len(svc.Middlewares) != 0,
			ProjectName:   layout.ProjectName,
			IdlName:       layout.IdlName,
			Pkg:           layout.Pkg,
		}

		for _, apiStr := range svc.Apis {
			api, err := getApi(layout, apiStr)
			if err != nil {
				return nil, err
			}
			api.ServiceName = svc.Name
			if m, ok := msg[api.ReqName]; ok {
				api.ReqParam = append(api.ReqParam, m.Param...)
			}

			group.HasMiddleware = group.HasMiddleware || api.HasMiddleware
			group.Apis = append(group.Apis, api)
		}
		res = append(res, group)
	}

	return res, nil
}

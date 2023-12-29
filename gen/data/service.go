package data

import (
	"go-server-gen/conf"
)

type Service struct {
	ServiceName   string                  // 当前service名称
	ProjectName   string                  // 当前项目名称，go mod name
	IdlName       string                  // idl name
	Pkg           map[string]conf.Package // 全局变量
	HasMiddleware bool                    // api或group是否包含中间件
	Middleware    []string                // group中间件
	Apis          []Api                   // 接口列表
	Handlers      map[string]string       // apis 解析后的结果
}

// ConfigToData 配置转数据
func ConfigToData(layout *conf.LayoutConfig, idl *conf.Idl) ([]Service, map[string]Message, error) {
	msg, err := getMessage(idl.Messages)
	if err != nil {
		return nil, nil, err
	}
	groups, err := getService(layout, idl.Services, msg)
	if err != nil {
		return nil, nil, err
	}
	return groups, msg, nil
}

// 解析service
func getService(layout *conf.LayoutConfig, services []conf.Service, msg map[string]Message) ([]Service, error) {
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

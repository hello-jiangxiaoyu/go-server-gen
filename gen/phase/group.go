package phase

import "go-server-gen/gen/conf"

type (
	Param struct {
		Name        string
		From        string
		Type        string
		Required    string
		Description string
	}
	Api struct {
		Method     string
		Path       string
		DocPath    string
		Summary    string
		FuncName   string
		ReqName    string
		Tag        string
		Context    string
		Middleware []string
		ReqParam   []Param
	}

	Group struct {
		Name          string
		ApiTemplate   string
		GroupTemplate string
		Import        []string
		Middleware    []string
		Apis          []Api
	}
)

func ConfigToCode(layout *conf.LayoutConfig, idl *conf.Idl) ([]Group, error) {
	msg, err := GetMessage(idl.Messages)
	if err != nil {
		return nil, err
	}
	return GenServiceCode(layout, idl.Services, msg)
}

func GenServiceCode(layoutConfig *conf.LayoutConfig, services []conf.Service, msg map[string]Message) ([]Group, error) {
	res := make([]Group, 0)
	for _, svc := range services {
		group := Group{
			Name:        svc.Name,
			Middleware:  svc.Middlewares,
			Import:      make([]string, 0),
			Apis:        make([]Api, 0),
			ApiTemplate: conf.DefaultApiTemplate,
		}

		for _, apiStr := range svc.Apis {
			api, err := getApi(layoutConfig, apiStr)
			if err != nil {
				return nil, err
			}
			api.Tag = svc.Name
			m, ok := msg[api.ReqName]
			if ok {
				api.ReqParam = append(api.ReqParam, m.Param...)
			}
			group.Apis = append(group.Apis, api)
		}
		res = append(res, group)
	}

	return res, nil
}

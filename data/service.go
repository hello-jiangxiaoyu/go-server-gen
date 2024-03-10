package data

import (
	"errors"
	"go-server-gen/conf"
	"go-server-gen/utils"
	"regexp"
	"strings"
)

// 解析service
func getService(layout conf.LayoutConfig, services []conf.Service, msg map[string]Message) ([]Service, error) {
	res := make([]Service, 0)
	for _, svc := range services {
		service := Service{
			ServiceName:   svc.Name,
			Middlewares:   svc.Middlewares,
			Apis:          make([]Api, 0),
			HasMiddleware: len(svc.Middlewares) != 0,
			ProjectName:   layout.ProjectName,
			IdlName:       layout.IdlName,
			Pkg:           layout.Pkg,
		}

		for _, apiStr := range svc.Apis {
			api, err := getApi(layout, apiStr)
			if err != nil {
				utils.Log("get api err: ", err.Error(), api)
				return nil, err
			}
			api.ServiceName = svc.Name
			if m, ok := msg[api.ReqName]; ok {
				api.ReqParam = append(api.ReqParam, m.Param...)
			}

			service.HasMiddleware = service.HasMiddleware || api.HasMiddleware
			service.Apis = append(service.Apis, api)
		}
		res = append(res, service)
	}

	return res, nil
}

var (
	regApi    = regexp.MustCompile(`(\w+)\("(.+?)"(.*)\)\s*//\s*(.+)`) // GET("/api/login", GetAppList)  // LoginReq
	regDocUri = regexp.MustCompile(`:(\w+)`)
)

// 解析api字符串
func getApi(layout conf.LayoutConfig, obj string) (Api, error) {
	matches := regApi.FindStringSubmatch(utils.RemoveSpace(obj))
	if len(matches) < 5 {
		return Api{}, errors.New("invalid API string")
	}
	res := Api{
		Method:   matches[1],
		Path:     matches[2],
		ReqName:  matches[4],
		IdlName:  layout.IdlName,
		Handlers: make([]string, 0),
		ReqParam: getUriParam(matches[2]),
	}
	handlers := strings.Split(matches[3], ",")
	if len(handlers) == 0 {
		return Api{}, errors.New("handler should not be empty")
	}

	var controllerFunc string
	handlers = utils.DeduplicateStrings(handlers)
	for _, handler := range handlers {
		res.Handlers = append(res.Handlers, handler)
		if strings.HasPrefix(handler, "middleware") {
			res.HasMiddleware = true
		} else {
			controllerFunc = handler
		}
	}
	res.FuncName = controllerFunc
	res.Summary = utils.ConvertToWord(controllerFunc, " ")
	res.Pkg = layout.Pkg
	res.ProjectName = layout.ProjectName
	res.IdlName = layout.IdlName
	return res, nil
}

// 提取路由中的参数
func getUriParam(path string) []Param {
	matches := regDocUri.FindAllStringSubmatch(path, -1)
	res := make([]Param, 0)
	for _, match := range matches {
		res = append(res, Param{
			Name:        match[1],
			Type:        getDocType(match[1]),
			From:        "path",
			Required:    "true",
			Description: utils.ConvertToWord(match[1], " "),
		})
	}
	return res
}

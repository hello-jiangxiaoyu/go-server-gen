package data

import (
	"errors"
	"go-server-gen/conf"
	"go-server-gen/utils"
	"regexp"
)

// 解析service
func getService(layout conf.LayoutConfig, services []conf.Service, msg map[string]Message) ([]Service, error) {
	res := make([]Service, 0)
	for _, svc := range services {
		service := Service{
			ServiceName: svc.Name,
			Apis:        make([]Api, 0),
			ProjectName: layout.ProjectName,
			IdlName:     layout.IdlName,
			Pkg:         layout.Pkg,
			MsgMap:      msg,
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

			service.Apis = append(service.Apis, api)
		}
		res = append(res, service)
	}

	return res, nil
}

var (
	regApi    = regexp.MustCompile(`(\w+)\("(.+?)",(.*)\)\s*//\s*(.+)`) // GET("/api/login", GetAppList)  // LoginReq
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
		Handler:  matches[3],
		FuncName: matches[3],
		ReqName:  matches[4],
		IdlName:  layout.IdlName,
		ReqParam: getUriParam(matches[2]),
	}
	res.Summary = utils.ConvertToWord(matches[3], " ")
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

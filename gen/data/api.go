package data

import (
	"errors"
	"go-server-gen/conf"
	"go-server-gen/utils"
	"regexp"
	"strings"
)

type Api struct {
	ServiceName   string                  // API所属service名
	Method        string                  // HTTP方法
	Path          string                  // 接口路径
	Summary       string                  // swagger 文档summary
	FuncName      string                  // 接口业务处理函数名
	ReqName       string                  // 请求参数名
	Handlers      []string                // 接口处理者，包含中间件
	ReqParam      []Param                 // 请求参数详情
	HasMiddleware bool                    // 是否包含中间件
	ProjectName   string                  // 当前项目名称，go mod name
	IdlName       string                  // idl name
	Pkg           map[string]conf.Package // layout定义的全局变量
}

var regApi = regexp.MustCompile(`(\w+)\("(.+?)"(.*)\)\s*//\s*(.+)`) // GET("/api/login", GetAppList)  // LoginReq

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
		Handlers: make([]string, 0),
		ReqParam: getUriParam(matches[2]),
	}
	handlers := strings.Split(matches[3], ",")
	if len(handlers) == 0 {
		return Api{}, errors.New("handler should not be empty: " + obj)
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

var regDocUri = regexp.MustCompile(`:(\w+)`)

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

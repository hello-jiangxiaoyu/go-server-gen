package phase

import (
	"errors"
	"go-server-gen/gen/conf"
	"go-server-gen/utils"
	"regexp"
	"strings"
)

var (
	regApi    = regexp.MustCompile(`(\w+)\("(.+?)"(.*)\)\s*//\s*(.+)`) // GET("/api/login", GetAppList)  // LoginReq
	regDocUri = regexp.MustCompile(`:(\w+)`)                           // 路由path转文档router
)

func getApi(layout *conf.LayoutConfig, apiStr string) (Api, error) {
	res, err := MatchApiString(apiStr)
	if err != nil {
		return Api{}, err
	}
	res.Context = layout.Server.ContextName
	return res, nil
}

func MatchApiString(obj string) (Api, error) {
	obj = strings.ReplaceAll(obj, " ", "")
	matches := regApi.FindStringSubmatch(obj)
	if len(matches) < 5 {
		return Api{}, errors.New("invalid API string")
	}
	res := Api{
		Method:     matches[1],
		Path:       matches[2],
		DocPath:    regDocUri.ReplaceAllString(matches[2], "{$1}"),
		ReqName:    matches[4],
		Middleware: make([]string, 0),
		ReqParam:   GetUriParam(matches[2]),
	}

	handlers := strings.Split(matches[3], ",")
	funcName := ""
	for _, handler := range handlers {
		if strings.HasPrefix(handler, "middleware") {
			res.Middleware = append(res.Middleware, handler)
		} else {
			funcName = handler
		}
	}
	res.FuncName = funcName
	res.Summary = utils.ConvertToWord(funcName, " ")
	return res, nil
}

func GetUriParam(path string) []Param {
	matches := regDocUri.FindAllStringSubmatch(path, -1)
	res := make([]Param, 0)
	for _, match := range matches {
		res = append(res, Param{
			Name:        match[1],
			Type:        GetDocType(match[1]),
			From:        "path",
			Required:    "true",
			Description: utils.ConvertToWord(match[1], " "),
		})
	}
	return res
}

package phase

import (
	"go-server-gen/gen/conf"
	"regexp"
	"strings"
)

type (
	Param struct {
		Name        string
		From        string
		Type        string
		Required    string
		Description string
	}
	Api struct {
		ServiceName   string
		Method        string
		Path          string
		Summary       string
		FuncName      string
		ReqName       string
		Handlers      []string
		ReqParam      []Param
		HasMiddleware bool
		ProjectName   string
		Pkg           map[string]conf.Package
	}

	Service struct {
		ServiceName   string
		ProjectName   string
		Pkg           map[string]conf.Package
		HasMiddleware bool     // api或group是否包含中间件
		Middleware    []string // group中间件
		Apis          []Api
		Handlers      map[string]string // apis 解析后的结果
	}
)

var (
	regQueryTag = regexp.MustCompile(`(?:query|formData):"([^"]+)"`) // 获取go结构体uri tag
)

func getRequestParam(msg string) (map[string]string, error) {
	matches := regQueryTag.FindAllStringSubmatch(msg, -1)
	res := make(map[string]string)
	for _, match := range matches {
		res[match[1]] = strings.Split(match[0], ":")[0]
	}

	return res, nil
}

func getDocType(name string) string {
	if strings.Contains(name, "ID") || strings.Contains(name, "Id") {
		return "integer"
	}
	if strings.Contains(name, "is") {
		return "bool"
	}
	return "string"
}

package utils

import (
	"regexp"
	"strings"
)

var (
	versionReg = regexp.MustCompile(`(^v\d+(\.\d+)*$)`)
	regDocUri  = regexp.MustCompile(`:(\w+)`) // 路由path转文档router
)

func Method(method string, ctxName string) string {
	if ctxName != "*fiber.Ctx" {
		return method
	}
	return UppercaseFirst(strings.ToLower(method))
}

func GetGoLastSplit(obj, sep string) string {
	res := strings.Split(obj, sep)
	if len(res) == 0 {
		return ""
	}
	if match := versionReg.FindStringSubmatch(res[len(res)-1]); len(match) > 0 && len(res) > 1 {
		return res[len(res)-2]
	}
	return res[len(res)-1]
}

func GetSwaggerRouter(path string) string {
	return regDocUri.ReplaceAllString(path, "{$1}")
}

// TS

func GetTsType(name string) string {
	if strings.HasSuffix(name, "ID") || strings.HasSuffix(name, "Id") {
		return "number"
	}
	if strings.HasPrefix(name, "is") {
		return "boolean"
	}
	return "string"
}

func GetTsRouter(path string) string {
	return regDocUri.ReplaceAllString(path, "${$1}")
}

func GetPathPara(path string) map[string]string {
	res := make(map[string]string)
	matches := regDocUri.FindAllString(path, -1)
	for i, _ := range matches {
		if len(matches[i]) > 0 {
			matches[i] = matches[i][1:]
		}
		res[matches[i]] = GetTsType(matches[i])
	}
	return res
}

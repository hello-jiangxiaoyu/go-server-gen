package utils

import (
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

var defaultFuncMap = template.FuncMap{
	"lowercaseFirst": LowercaseFirst,    // 首字母小写
	"uppercaseFirst": UppercaseFirst,    // 首字母小写
	"convertToWord":  ConvertToWord,     // 路由转文档Router
	"hasPrefix":      strings.HasPrefix, // 是否包含前缀
	"removeSpace":    RemoveSpace,       // 去除空格
	"getDocRouter":   GetDocRouter,      // 转文档Router
	"getTsRouter":    GetTsRouter,       // 转ts路径
	"getPathPara":    GetPathPara,       // 获取路径中的参数
	"getTsType":      GetTsType,         // 根据名字推测ts类型
	"join":           strings.Join,      // 切片
	"mapJoin":        MapJoin,           // map拼接
}

// LowercaseFirst 首字母小写
func LowercaseFirst(obj string) string {
	runes := []rune(obj)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// UppercaseFirst 首字母大写
func UppercaseFirst(obj string) string {
	runes := []rune(obj)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

var regDocUri = regexp.MustCompile(`:(\w+)`) // 路由path转文档router

func GetDocRouter(path string) string {
	return regDocUri.ReplaceAllString(path, "{$1}")
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

func MapJoin(obj map[string]string, mid, sep string, withEnd bool) string {
	res := ""
	for key, value := range obj {
		res += key + mid + value + sep
	}
	if len(res) > len(sep) && !withEnd {
		res = res[:len(res)-len(sep)]
	}
	return res
}
func GetTsType(name string) string {
	if strings.Contains(name, "ID") || strings.Contains(name, "Id") {
		return "number"
	}
	if strings.Contains(name, "is") {
		return "boolean"
	}
	return "string"
}

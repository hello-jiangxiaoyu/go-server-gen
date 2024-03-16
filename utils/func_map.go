package utils

import (
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

var defaultFuncMap = template.FuncMap{
	"hasPrefix":      strings.HasPrefix, // 是否包含前缀
	"hasSuffix":      strings.HasSuffix, // 是否包含后缀
	"join":           strings.Join,      // 切片
	"lowercaseFirst": LowercaseFirst,    // 首字母小写
	"uppercaseFirst": UppercaseFirst,    // 首字母大写
	"removeSpace":    RemoveSpace,       // 去除空格
	"convertToWord":  ConvertToWord,     // 字符串拆分拼接
	"getPathPara":    GetPathPara,       // 获取路径中的参数
	"getFirstSplit":  GetFirstSplit,     // 获取第一个单词
	"mapJoin":        MapJoin,           // map拼接

	"getDocRouter":   GetDocRouter,   // 转文档Router
	"getTsRouter":    GetTsRouter,    // 转ts路径
	"getTsType":      GetTsType,      // 根据名字推测ts类型
	"getGoLastSplit": GetGoLastSplit, // 获取最后一个单词
	"method":         Method,         // method方法特殊处理
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

func Method(method string, ctxName string) string {
	if ctxName != "*fiber.Ctx" {
		return method
	}
	return UppercaseFirst(strings.ToLower(method))
}

var versionReg = regexp.MustCompile(`(^v\d+(\.\d+)*$)`)

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

func GetFirstSplit(obj, sep string) string {
	res := strings.Split(obj, sep)
	return res[0]
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
	if strings.HasSuffix(name, "ID") || strings.HasSuffix(name, "Id") {
		return "number"
	}
	if strings.HasPrefix(name, "is") {
		return "boolean"
	}
	return "string"
}

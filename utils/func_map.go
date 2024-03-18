package utils

import (
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

	"convertToWord": ConvertToWord, // 字符串拆分拼接
	"getFirstSplit": GetFirstSplit, // 获取第一个单词
	"mapJoin":       MapJoin,       // map拼接
	"getPathPara":   GetPathPara,   // 获取路径中的参数

	"getDocRouter":   GetSwaggerRouter, // 转文档Router
	"getGoLastSplit": GetGoLastSplit,   // 获取go import最后一个单词，排除version
	"method":         Method,           // method方法特殊处理
	"getTsRouter":    GetTsRouter,      // 转ts路径
	"getTsType":      GetTsType,        // 根据名字推测ts类型
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

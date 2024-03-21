package template

import (
	"go-server-gen/utils"
	"strings"
	"text/template"
)

var defaultFuncMap = template.FuncMap{
	"hasPrefix": strings.HasPrefix, // 是否包含前缀
	"hasSuffix": strings.HasSuffix, // 是否包含后缀
	"contains":  strings.Contains,  // 包含
	"join":      strings.Join,      // 切片

	"lowercaseFirst":        utils.LowercaseFirst,        // 首字母小写
	"uppercaseFirst":        utils.UppercaseFirst,        // 首字母大写
	"snakeToUpperCamelCase": utils.SnakeToUpperCamelCase, // 蛇形转大驼峰
	"snakeToLowerCamelCase": utils.SnakeToLowerCamelCase, // 蛇形转小驼峰

	"removeSpace":   utils.RemoveSpace,   // 去除空格
	"convertToWord": utils.ConvertToWord, // 字符串拆分拼接
	"getFirstSplit": utils.GetFirstSplit, // 获取第一个单词
	"mapJoin":       utils.MapJoin,       // map拼接

	"removeSuffix": RemoveSuffix,

	"getPathPara":    GetPathPara,      // 获取路径中的参数
	"getDocRouter":   GetSwaggerRouter, // 转文档Router
	"getGoLastSplit": GetGoLastSplit,   // 获取go import最后一个单词，排除version
	"getGoMethod":    GetGoMethod,      // Go method方法特殊处理
	"getTsRouter":    GetTsRouter,      // 转ts路径
	"getTsType":      GetTsType,        // 根据名字推测ts类型
}

func RemoveSuffix(s string, suf string) string {
	if len(suf) > 0 && strings.HasSuffix(s, suf) {
		return s[:len(s)-len(suf)]
	}
	return s
}
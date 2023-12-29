package utils

import (
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

var defaultFuncMap = template.FuncMap{
	"uncapitalizeFirstLetter": UncapitalizeFirstLetter, // 首字母小写
	"getDocRouter":            GetDocRouter,            // 路由转文档Router
	"convertToWord":           ConvertToWord,           // 路由转文档Router
	"hasPrefix":               strings.HasPrefix,       // 是否包含前缀
	"removeSpace":             RemoveSpace,             // 去除空格
}

// UncapitalizeFirstLetter 首字母小写
func UncapitalizeFirstLetter(obj string) string {
	runes := []rune(obj)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

var regDocUri = regexp.MustCompile(`:(\w+)`) // 路由path转文档router

func GetDocRouter(path string) string {
	return regDocUri.ReplaceAllString(path, "{$1}")
}

// ConvertToWord 驼峰字符串拆分成句子
func ConvertToWord(obj string, mid string) string {
	words := SplitCamelCase(obj)
	content := ""
	idFlag := false // ID不拆分
	for _, word := range words {
		if word == "d" && idFlag {
			content += word
		} else {
			content += mid + word
		}
		idFlag = word == "i"
	}
	if content != "" {
		content = content[1:]
	}
	return content
}

// RemoveSpace 删除字符串中的空格
func RemoveSpace(obj string) string {
	return strings.ReplaceAll(obj, " ", "")
}

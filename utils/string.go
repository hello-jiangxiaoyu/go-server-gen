package utils

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

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

// SnakeToLowerCamelCase 小驼峰
func SnakeToLowerCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		r, size := utf8.DecodeRuneInString(parts[i])
		parts[i] = string(unicode.ToLower(r)) + parts[i][size:]
	}

	return strings.Join(parts, "")
}

// SnakeToUpperCamelCase 大驼峰
func SnakeToUpperCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 0; i < len(parts); i++ {
		r, size := utf8.DecodeRuneInString(parts[i])
		parts[i] = string(unicode.ToLower(r)) + parts[i][size:]
	}

	return strings.Join(parts, "")
}

// RemoveSpace 删除字符串中的空格
func RemoveSpace(obj string) string {
	return strings.ReplaceAll(obj, " ", "")
}

func GetFirstSplit(obj, sep string) string {
	res := strings.Split(obj, sep)
	return res[0]
}

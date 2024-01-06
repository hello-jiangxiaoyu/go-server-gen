package utils

import (
	"regexp"
	"strings"
)

// SplitCamelCase 驼峰字符串拆分成单词
func SplitCamelCase(str string) []string {
	re := regexp.MustCompile(`[A-Za-z][^A-Z]*`)
	words := re.FindAllString(str, -1)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
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

// DeduplicateStrings 字符串数组去重并去除空字符串
func DeduplicateStrings(arr []string) []string {
	visited := make(map[string]bool)
	result := make([]string, 0)
	for _, str := range arr {
		if str == "" {
			continue
		}
		if !visited[str] {
			visited[str] = true
			result = append(result, str)
		}
	}

	return result
}

// RemoveSpace 删除字符串中的空格
func RemoveSpace(obj string) string {
	return strings.ReplaceAll(obj, " ", "")
}

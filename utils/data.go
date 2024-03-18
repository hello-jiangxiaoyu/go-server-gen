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

package utils

import (
	"regexp"
	"strings"
	"unicode"
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
	sp := SplitCamelCase(obj)
	res := ""
	idFlag := false
	for _, v := range sp {
		if v == "d" && idFlag {
			res += v
		} else {
			res += mid + v
		}
		if v == "i" {
			idFlag = true
		} else {
			idFlag = false
		}
	}
	if res != "" {
		res = res[1:]
	}

	return res
}

// UncapitalizeFirstLetter 首字母小写
func UncapitalizeFirstLetter(obj string) string {
	runes := []rune(obj)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
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

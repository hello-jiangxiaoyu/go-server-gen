package utils

import (
	"encoding/json"
	"errors"
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

func StructToString(obj any) string {
	res, err := make([]byte, 0), errors.New("")
	if res, err = json.Marshal(&obj); err != nil {
		return ""
	}
	return string(res)
}

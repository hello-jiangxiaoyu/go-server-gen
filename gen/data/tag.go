package data

import (
	"regexp"
	"strings"
)

var (
	regQueryTag = regexp.MustCompile(`(?:query|formData):"([^"]+)"`) // 获取go结构体uri tag
)

func getRequestParam(msg string) (map[string]string, error) {
	matches := regQueryTag.FindAllStringSubmatch(msg, -1)
	res := make(map[string]string)
	for _, match := range matches {
		res[match[1]] = strings.Split(match[0], ":")[0]
	}

	return res, nil
}

func getDocType(name string) string {
	if strings.Contains(name, "ID") || strings.Contains(name, "Id") {
		return "integer"
	}
	if strings.Contains(name, "is") {
		return "bool"
	}
	return "string"
}

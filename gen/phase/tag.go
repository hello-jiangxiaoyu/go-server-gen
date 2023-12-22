package phase

import (
	"regexp"
	"strings"
)

var (
	regQueryTag = regexp.MustCompile(`query:"([^"]+)"`) // 获取go结构体uri tag
)

func GetQueryPara(msg string) ([]string, error) {
	matches := regQueryTag.FindAllStringSubmatch(msg, -1)
	res := make([]string, 0)
	for _, match := range matches {
		res = append(res, match[1])
	}

	return res, nil
}

func GetDocType(name string) string {
	if strings.Contains(name, "ID") || strings.Contains(name, "Id") {
		return "integer"
	}
	if strings.Contains(name, "is") {
		return "bool"
	}
	return "string"
}

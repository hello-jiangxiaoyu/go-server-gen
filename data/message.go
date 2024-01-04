package data

import (
	"fmt"
	"go-server-gen/utils"
	"go/format"
	"regexp"
	"strings"
)

var (
	regMessageSplit = regexp.MustCompile(`type (\w+) struct {([^}]+)}`)
	regQueryTag     = regexp.MustCompile(`(?:query|formData):"([^"]+)"`) // 获取go结构体uri tag
)

func getMessage(msg string) (map[string]Message, error) {
	res := make(map[string]Message)
	messages, err := splitMessage(msg)
	if err != nil {
		return nil, err
	}
	for structName, structBody := range messages {
		queryMatches := regQueryTag.FindAllStringSubmatch(structBody, -1)
		queries := make(map[string]string)
		for _, match := range queryMatches {
			queries[match[1]] = strings.Split(match[0], ":")[0]
		}

		param := make([]Param, 0)
		for queryName, queryType := range queries {
			param = append(param, Param{
				Name:        queryName,
				From:        queryType,
				Type:        getDocType(queryName),
				Required:    "false",
				Description: queryName,
			})
		}
		res[structName] = Message{
			Name:   structName,
			Param:  param,
			Source: structBody,
		}
	}

	return res, nil
}

func splitMessage(msgCode string) (map[string]string, error) {
	code, err := format.Source([]byte(msgCode))
	if err != nil {
		utils.Log("split message err: ", err.Error())
		return nil, err
	}

	res := make(map[string]string)
	matches := regMessageSplit.FindAllStringSubmatch(string(code), -1)
	for _, match := range matches {
		typeName := match[1]
		typeBody := strings.TrimSpace(match[2])
		tmpCode := fmt.Sprintf("type %s struct {\n\t%s\n}\n", typeName, typeBody)
		res[typeName] = tmpCode
	}
	return res, nil
}

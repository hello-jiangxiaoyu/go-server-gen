package data

import (
	"fmt"
	"go-server-gen/conf"
	"go-server-gen/utils"
	"go/format"
	"regexp"
	"strings"
)

var (
	regMessageSplit = regexp.MustCompile(`type (\w+) struct {([^}]+)}`)
	regQueryTag     = regexp.MustCompile(`(?:query|formData):"([^"]+)"`) // 获取go结构体uri tag
)

func getMessage(idl conf.IdlConfig) (map[string]Message, error) {
	res := make(map[string]Message)
	messages, err := splitGoMessage(idl.Messages)
	if err != nil {
		return nil, err
	}
	for structName, structBody := range messages {
		res[structName] = Message{
			Name:   structName,
			Lang:   "go",
			Source: structBody,
		}
	}
	res["__ts"] = Message{
		Name:   "__ts",
		Lang:   "ts",
		Source: idl.Ts,
	}

	return res, nil
}

func splitGoMessage(msgCode string) (map[string]string, error) {
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

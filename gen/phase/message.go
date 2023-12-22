package phase

import (
	"fmt"
	"go-server-gen/utils"
	"regexp"
	"strings"
)

var (
	regMessageSplit = regexp.MustCompile(`type (\w+) struct {([^}]+)}`)
)

type Message struct {
	Param  []Param `json:"param"`
	Source string  `json:"source"`
}

func GetMessage(msg string) (map[string]Message, error) {
	res := make(map[string]Message)
	sp, err := SplitMessage(msg)
	if err != nil {
		return nil, err
	}
	for k, v := range sp {
		param := make([]Param, 0)
		query, err := GetQueryPara(v)
		if err != nil {
			return nil, err
		}
		for _, q := range query {
			param = append(param, Param{
				Name:        q,
				From:        "query",
				Type:        GetDocType(q),
				Required:    "false",
				Description: q,
			})
		}
		res[k] = Message{
			Param:  param,
			Source: v,
		}
	}

	return res, nil
}

func SplitMessage(msgCode string) (map[string]string, error) {
	code, err := utils.FormatCode([]byte(msgCode))
	if err != nil {
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

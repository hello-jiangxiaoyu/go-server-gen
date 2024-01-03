package data

import (
	"go-server-gen/conf"
	"strings"
)

// ConfigToData 配置转数据
func ConfigToData(layout conf.LayoutConfig, idl conf.Idl) ([]Service, map[string]Message, error) {
	msg, err := getMessage(idl.Messages)
	if err != nil {
		return nil, nil, err
	}
	groups, err := getService(layout, idl.Services, msg)
	if err != nil {
		return nil, nil, err
	}
	return groups, msg, nil
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

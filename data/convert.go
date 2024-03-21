package data

import (
	"go-server-gen/conf"
	"strings"
)

// ConfigToData 配置转数据
func ConfigToData(layout conf.LayoutConfig, idl conf.IdlConfig) ([]Service, error) {
	msg, err := getMessage(idl)
	if err != nil {
		return nil, err
	}
	groups, err := GetServiceData(layout, idl.Services, msg)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// 根据变量名推测变量类型
func getDocType(name string) string {
	if strings.Contains(name, "ID") || strings.Contains(name, "Id") {
		return "integer"
	}
	if strings.Contains(name, "is") {
		return "bool"
	}
	return "string"
}

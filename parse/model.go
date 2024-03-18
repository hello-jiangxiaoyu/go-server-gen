package parse

import (
	_ "embed"
	"fmt"
	"go-server-gen/conf"
	"go-server-gen/utils"
)

type ViewColumn struct {
	Key           string   `json:"key"`
	Column        string   `json:"column"`
	Label         string   `json:"label"`
	LabelWidth    int      `json:"labelWidth"`
	Type          string   `json:"type"`
	ViewType      string   `json:"viewType"` // 前端展示类型
	CanCreate     bool     `json:"canCreate"`
	CanEdit       bool     `json:"canEdit"`
	CanSearch     bool     `json:"canSearch"`
	Required      bool     `json:"required"`
	Placeholder   string   `json:"placeholder"`
	SelectOptions []string `json:"selectOptions"`
}

func GetIdlConfig(tableName string, columns []ViewColumn) (conf.IdlConfig, error) {
	service := conf.Service{
		Name: tableName,
		Apis: make([]string, 0),
	}
	for _, api := range conf.ConstData.Api {
		service.Apis = append(service.Apis, api)
	}

	return conf.IdlConfig{
		Ts:       getTsArrayObject(columns),
		Messages: getGoStructRows(tableName, columns),
		Services: []conf.Service{service},
	}, nil
}

func getGoStructRows(tableName string, columns []ViewColumn) string {
	res := fmt.Sprintf("type %s struct{", utils.SnakeToUpperCamelCase(tableName))
	for _, v := range conf.ConstData.GoStruct {
		res += "\n" + v
	}
	for _, v := range columns {
		row := fmt.Sprintf("\n%s %s `%s`", UppercaseFirst(v.Column), getGoType(v.ViewType), getJson(v.Column))
		res += row
	}
	return res + "\n}"
}

func getJson(column string) string {
	return fmt.Sprintf(`json:"%s"`, utils.SnakeToLowerCamelCase(column))
}

func getGoType(viewType string) string {
	switch viewType {
	case "number":
		return "int"
	case "switch":
		return "bool"
	case "date", "time", "datetime":
		return "time.Time"
	}

	return "string"
}

func getTsArrayObject(columns []ViewColumn) string {
	res := ""
	for _, v := range columns {
		res += "{" + getTsColumns(v) + "\n},"
	}
	return res
}

func getTsColumns(column ViewColumn) string {
	res := "\ncolumn: '" + column.Column + "',"
	res += "\nlabel: '" + column.Label + "',"

	if column.CanEdit {
		res += "\neditable: true,"
	}
	if column.CanSearch {
		res += "\nsearchable: true,"
	}
	if column.CanCreate {
		res += "\ncanCreate: true,"
	}
	if column.Required {
		res += "\nrequire: true,"
	}

	if len(column.Placeholder) > 0 {
		res += "\nplaceholder: '" + column.Placeholder + "',"
	}

	return res
}

package parse

import (
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

func GetIdlConfig(tableName string, columns []ViewColumn) conf.Idl {
	service := conf.Service{
		Name:        tableName,
		Middlewares: make([]string, 0),
		Apis:        make([]string, 0),
	}
	for _, api := range apiMap {
		service.Apis = append(service.Apis, api)
	}

	_ = getGoFieldMap(columns)
	msg := ""
	return conf.Idl{
		Messages: msg,
		Ts:       "export interface FormStruct [" + convertToTs(columns) + "]",
		Services: []conf.Service{service},
	}
}

var apiMap = map[string]string{
	"get":         `GET("{{.Prefix}}/{{lowercaseFirst .ServiceName}}", Get{{uppercaseFirst .ServiceName}}List)  // {{uppercaseFirst .ServiceName}}Req`,
	"list":        `GET("{{.Prefix}}/{{lowercaseFirst .ServiceName}}/:{{lowercaseFirst .ServiceName}}ID", Get{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
	"create":      `POST("{{.Prefix}}/{{lowercaseFirst .ServiceName}}", Create{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
	"update":      `PUT("{{.Prefix}}/{{lowercaseFirst .ServiceName}}/:{{lowercaseFirst .ServiceName}}ID", Update{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
	"delete":      `DELETE("{{.Prefix}}/{{lowercaseFirst .ServiceName}}/:{{lowercaseFirst .ServiceName}}ID", Delete{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
	"batchDelete": `DELETE("{{.Prefix}}/{{lowercaseFirst .ServiceName}}", BatchDelete{{uppercaseFirst .ServiceName}})  // {{uppercaseFirst .ServiceName}}Req`,
}

func convertToTs(columns []ViewColumn) string {
	var res string
	for _, v := range columns {
		res += "\n{" + getTsColumns(v) + "\n},"
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

func getGoFieldMap(columns []ViewColumn) map[string]string {
	res := map[string]string{
		"All":    "",
		"Edit":   "",
		"Create": "",
		"Search": "",
	}
	for _, v := range columns {
		row := fmt.Sprintf("\n%s %s `%s`", utils.UppercaseFirst(v.Column), getGoType(v.ViewType), getJson(v.Column))
		res["All"] += row
		if v.CanEdit {
			res["Edit"] += row
		}
		if v.CanCreate {
			res["Create"] += row
		}
		if v.CanSearch {
			res["Search"] += row
		}
	}

	return res
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

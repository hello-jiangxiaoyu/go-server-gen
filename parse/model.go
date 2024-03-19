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

type GenRequest struct {
	TableName    string          `json:"-"`
	ProjectName  string          `json:"-"`
	RouterPrefix string          `json:"-"`
	Apis         map[string]bool `json:"apis"`
	Columns      []ViewColumn    `json:"columns"`
}

func GetIdlConfig(req GenRequest) (conf.IdlConfig, error) {
	service := conf.Service{
		Name: req.TableName,
		Apis: make([]string, 0),
	}
	for apiName, open := range req.Apis {
		if open {
			if res, ok := conf.ConstData.Api[apiName]; ok {
				api, err := ParseTemplate(res, map[string]any{
					"ProjectName": req.ProjectName,
					"ServiceName": req.TableName,
					"Prefix":      req.RouterPrefix,
				})
				if err != nil {
					return conf.IdlConfig{}, err
				}
				service.Apis = append(service.Apis, api)
			}
		}
	}

	ts, err := ParseTemplate(getTsArrayObject(req.Columns), map[string]any{
		"ProjectName": req.ProjectName,
		"ServiceName": req.TableName,
		"Prefix":      req.RouterPrefix,
	})
	if err != nil {
		return conf.IdlConfig{}, err
	}
	msg, err := ParseTemplate(getGoStructRows(req.TableName, req.Columns), map[string]any{
		"ProjectName": req.ProjectName,
		"ServiceName": req.TableName,
		"Prefix":      req.RouterPrefix,
	})
	return conf.IdlConfig{
		Ts:       ts,
		Messages: msg,
		Services: []conf.Service{service},
	}, nil
}

func getGoStructRows(tableName string, columns []ViewColumn) string {
	res := fmt.Sprintf("type %s struct{", utils.UppercaseFirst(tableName))
	for _, v := range conf.ConstData.GoStruct {
		res += "\n" + v
	}
	for _, v := range columns {
		row := fmt.Sprintf("\n%s %s `%s`", utils.SnakeToUpperCamelCase(v.Column), getGoType(v.ViewType), getJson(v.Column))
		if v.CanEdit || v.CanCreate || v.CanSearch {
			res += row
		}
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

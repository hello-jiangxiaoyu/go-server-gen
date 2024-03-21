package data

import (
	"fmt"
	"go-server-gen/conf"
	"go-server-gen/template"
	"go-server-gen/utils"
)

type GenRequest struct {
	TableName    string          `json:"-"`
	ProjectName  string          `json:"-"`
	RouterPrefix string          `json:"routerPrefix"`
	Apis         map[string]bool `json:"apis"`
	Columns      []ViewColumn    `json:"columns"`
}

func GenServiceByTable(req GenRequest, layout conf.LayoutConfig) (Service, error) {
	msg, err := genGoStruct(req)
	if err != nil {
		return Service{}, utils.WithMessage(err, "gen go struct err")
	}

	modelName := utils.UppercaseFirst(req.TableName)
	service := Service{
		ServiceName: req.TableName,
		Apis:        make([]Api, 0),
		ProjectName: layout.ProjectName,
		IdlName:     layout.IdlName,
		Pkg:         layout.Pkg,
		MsgMap: map[string]Message{
			modelName:       msg,
			"__tsInterface": {Source: genTsTableInterface(req), Lang: "ts"},
			"__tsStruct":    {Source: genTsTableStruct(req), Lang: "ts"},
		},
	}

	for k, v := range conf.ConstData.Api {
		if apiSwitch, ok := req.Apis[k]; !ok || !apiSwitch {
			continue
		}
		apiStr, err := template.ParseTemplate(v, map[string]any{
			"ProjectName": req.ProjectName,
			"ServiceName": req.TableName,
			"Prefix":      req.RouterPrefix,
		})
		if err != nil {
			return Service{}, utils.WithMessage(err, "api str template err")
		}
		api, err := getApi(layout, apiStr)
		if err != nil {
			return Service{}, utils.WithMessage(err, "get init api err")
		}

		api.ServiceName = req.TableName
		api.Msg = msg
		api.MsgMap = service.MsgMap
		api.BodyColumns = getApiBodyColumns(k, req)
		if k == "list" {
			api.ReqParam = append(api.ReqParam,
				Param{Name: "pageNumber", From: "query", Type: "integer", Required: "false", Description: "page number"},
				Param{Name: "pageSize", From: "query", Type: "integer", Required: "false", Description: "page size"},
			)
		} else if k == "create" || k == "update" || k == "batchDelete" {
			api.ReqParam = append(api.ReqParam, Param{Name: "data", From: "body", Type: "models." + modelName, Required: "true", Description: "body"})
		}
		service.Apis = append(service.Apis, api)
	}

	return service, nil
}

func getApiBodyColumns(key string, req GenRequest) []ViewColumn {
	res := make([]ViewColumn, 0)
	switch key {
	case "create":
		for _, c := range req.Columns {
			if c.CanCreate {
				res = append(res, c)
			}
		}
	case "update":
		for _, c := range req.Columns {
			if c.CanEdit {
				res = append(res, c)
			}
		}
	case "batchDelete":
		res = append(res, ViewColumn{
			Column: utils.SnakeToUpperCamelCase(req.TableName) + "IDs",
			Type:   "[]int64",
		})
	}
	return res
}

func genGoStruct(req GenRequest) (Message, error) {
	res := fmt.Sprintf("type %s struct{", utils.UppercaseFirst(req.TableName))
	for _, v := range conf.ConstData.GoStruct {
		res += "\n\t" + v
	}
	for _, v := range req.Columns {
		jsonField := fmt.Sprintf(`json:"%s,omitempty"`, utils.SnakeToLowerCamelCase(v.Column))
		row := fmt.Sprintf("\n\t%s %s `%s`", utils.SnakeToUpperCamelCase(v.Column), v.GoType, jsonField)
		if v.CanEdit || v.CanCreate || v.CanSearch {
			res += row
		}
	}
	goSource, err := template.ParseTemplate(res+"\n}", map[string]any{
		"ProjectName": req.ProjectName,
		"ServiceName": req.TableName,
		"Prefix":      req.RouterPrefix,
	})
	if err != nil {
		return Message{}, err
	}

	return Message{
		Lang:    "go",
		Source:  goSource,
		Columns: req.Columns,
	}, nil
}

func genTsTableStruct(req GenRequest) string {
	res := fmt.Sprintf("const %sTableStruct: Array<ItfTableStruct<%s>> = [\n", utils.SnakeToLowerCamelCase(req.TableName), utils.SnakeToUpperCamelCase(req.TableName))
	for _, v := range req.Columns {
		tsColumn := fmt.Sprintf("\n    column: '%s',", v.Column)
		tsColumn += fmt.Sprintf("\n    label: '%s',", v.Label)
		if v.CanEdit {
			tsColumn += "\n    editable: true,"
		}
		if v.CanSearch {
			tsColumn += "\n    searchable: true,"
		}
		if v.CanCreate {
			tsColumn += "\n    canCreate: true,"
		}
		if v.Required {
			tsColumn += "\n    require: true,"
		}
		if len(v.ViewType) > 0 {
			tsColumn += fmt.Sprintf("\n    viewType: '%s',", v.ViewType)
		}
		if v.LabelWidth > 0 {
			tsColumn += fmt.Sprintf("\n    labelWidth: %d,", v.LabelWidth)
		}
		if len(v.Placeholder) > 0 {
			tsColumn += fmt.Sprintf("\n    placeholder: '%s',", v.Placeholder)
		}
		res += "  {" + tsColumn + "\n  },\n"
	}
	return res + "]"
}

func genTsTableInterface(req GenRequest) string {
	res := fmt.Sprintf("export interface %s {", utils.SnakeToUpperCamelCase(req.TableName))
	for _, v := range req.Columns {
		if !v.CanEdit && !v.CanSearch && !v.CanCreate {
			continue
		}
		res += fmt.Sprintf("\n  %s: %s", utils.SnakeToLowerCamelCase(v.Column), v.TsType)
	}
	return res + "\n}"
}

package cmd

import (
	_ "embed"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"os"

	"github.com/spf13/cobra"
)

//go:embed crud-tpl.yaml
var CrudTemplate string

func CreateCrudGroup(_ *cobra.Command, _ []string) {
	if CrudServiceName == "" {
		println("service name is empty")
		os.Exit(1)
	}
	projectName, err := utils.GetProjectName()
	if err != nil {
		println("Failed to get project name: ", err.Error())
		os.Exit(1)
	}
	body, err := utils.ParseTemplate(CrudTemplate, map[string]any{
		"ProjectName": projectName,
		"ServiceName": CrudServiceName,
		"Prefix":      CrudRouterPrefix,
	})
	if err != nil {
		println("Failed to parse crud template: ", err.Error())
		os.Exit(1)
	}

	path := utils.ConvertToWord(CrudServiceName, "-") + ".yaml"
	err = writer.Write(map[string]writer.WriteCode{
		path: {
			File:  path,
			Code:  body,
			Write: writer.Skip,
		},
	}, "")
	if err != nil {
		println("Failed to write curd config file: ", err.Error())
		os.Exit(1)
	}
	println("Success")
}

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

func CreateCrudGroup(_ *cobra.Command, args []string) {
	checkCreateCmdArgs(args)

	projectName, err := utils.GetProjectName()
	if err != nil {
		utils.Log("Failed to get project name: ", err.Error())
		os.Exit(1)
	}
	path := utils.ConvertToWord(CrudServiceName, "-") + ".yaml"
	if utils.FileExists(path) && !ForceWrite {
		utils.Log(path + " is already exists! use --force to overwrite")
		os.Exit(1)
	}

	body, err := utils.ParseTemplate(CrudTemplate, map[string]any{
		"ProjectName": projectName,
		"ServiceName": CrudServiceName,
		"Prefix":      CrudRouterPrefix,
	})
	if err != nil {
		utils.Log("Failed to parse crud template: ", err.Error())
		os.Exit(1)
	}

	writeType := writer.Overwrite
	if !ForceWrite {
		writeType = writer.Skip
	}
	if err = writer.Write(map[string]writer.WriteCode{
		path: {
			File:  path,
			Code:  body,
			Write: writeType,
		},
	}); err != nil {
		utils.Log("Failed to write curd config file: ", err.Error())
		os.Exit(1)
	}

	println("Success")
}

func checkCreateCmdArgs(args []string) {
	if len(args) == 0 || args[0] == "" {
		utils.Log("service name should not be empty")
		os.Exit(1)
	}
	CrudServiceName = args[0]
}

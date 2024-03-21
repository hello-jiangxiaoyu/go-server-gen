package cmd

import (
	_ "embed"
	"go-server-gen/template"
	"go-server-gen/utils"
	"go-server-gen/writer"
	"os"
	"strings"

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

	body, err := template.ParseTemplate(CrudTemplate, map[string]any{
		"ProjectName": projectName,
		"ServiceName": CrudServiceName,
		"Prefix":      RouterPrefix,
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
	goMod, err := os.ReadFile("go.mod")
	if err != nil {
		utils.Log("failed to read go.mod: ", err.Error())
		os.Exit(1)
	}

	if ServerType == "" {
		if strings.Contains(string(goMod), "github.com/gin-gonic/gin") {
			ServerType = "gin"
		} else if strings.Contains(string(goMod), "github.com/gofiber/fiber") {
			ServerType = "fiber"
		} else if strings.Contains(string(goMod), "github.com/labstack/echo") {
			ServerType = "echo"
		} else if strings.Contains(string(goMod), "github.com/cloudwego/hertz") {
			ServerType = "hertz"
		}
	}
	if ServerType == "" && LayoutPath == "" {
		utils.Log("server type is empty")
		os.Exit(1)
	}
}

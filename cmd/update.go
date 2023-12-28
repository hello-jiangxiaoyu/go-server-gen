package cmd

import (
	"github.com/spf13/cobra"
	"go-server-gen/gen"
	"go-server-gen/utils"
	"os"
)

func UpdateProject(_ *cobra.Command, _ []string) {
	if _, err := utils.GetProjectName(); err != nil {
		println("failed to get project name by execute `go list`: " + err.Error())
		os.Exit(1)
	}
	if err := gen.InitConfig(LayoutPath, IdlPath); err != nil {
		println("init config err: ", err.Error())
		os.Exit(1)
	}
	if err := gen.Execute(); err != nil {
		println("exec err: ", err.Error())
		os.Exit(1)
	}
}

package cmd

import (
	_ "embed"
	"github.com/spf13/cobra"
	"go-server-gen/conf"
	"os"
)

func NewProject(_ *cobra.Command, _ []string) {
	if err := conf.InitConfig(LayoutPath, IdlPath); err != nil {
		println("init config err: ", err.Error())
		os.Exit(1)
	}
	if IdlPath != "" {
		CreateProjectTpl = "nil" // 指定idl时，模板无效
	}
}

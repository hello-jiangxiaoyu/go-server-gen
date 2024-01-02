package cmd

import (
	"go-server-gen/conf"
	"go-server-gen/gen"
	"os"

	"github.com/spf13/cobra"
)

func CreateCrudGroup(_ *cobra.Command, _ []string) {
	if err := conf.InitConfig(LayoutPath, IdlPath); err != nil {
		println("init config err: ", err.Error())
		os.Exit(1)
	}
	if err := gen.ExecuteUpdate(); err != nil {
		println("exec err: ", err.Error)
		os.Exit(1)
	}
}

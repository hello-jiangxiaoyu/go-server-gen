package cmd

import (
	"github.com/spf13/cobra"
	"go-server-gen/conf"
	"go-server-gen/gen"
	"os"
)

func UpdateProject(_ *cobra.Command, _ []string) {
	if err := conf.InitConfig(LayoutPath, IdlPath); err != nil {
		println("init config err: ", err.Error())
		os.Exit(1)
	}
	if err := gen.ExecuteUpdate("hertz"); err != nil {
		println("exec err: ", err.Error())
		os.Exit(1)
	}

	// gsg update --server gin --idl test-idl.yaml --ts true
}

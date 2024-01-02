package cmd

import (
	"github.com/spf13/cobra"
	"go-server-gen/gen"
	"os"
)

func UpdateProject(_ *cobra.Command, _ []string) {
	if IdlPath == "" {
		println("idl path is empty")
		os.Exit(1)
	}

	if err := gen.ExecuteUpdate(ServerType, LayoutPath, IdlPath); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	println("Success")
}

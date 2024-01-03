package cmd

import (
	"go-server-gen/gen"
	"os"

	"github.com/spf13/cobra"
)

func CreateCrudGroup(_ *cobra.Command, _ []string) {
	if err := gen.ExecuteUpdate(ServerType, LayoutPath, IdlPath, ""); err != nil {
		println("exec err: ", err.Error)
		os.Exit(1)
	}
}

package main

import (
	"github.com/spf13/cobra"
	"go-server-gen/cmd"
	"os"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "gen",
		Run: cmd.UpdateProject,
	}
	rootCmd.AddCommand(&cobra.Command{
		Use: "new",
		Run: cmd.NewProject,
	}, &cobra.Command{
		Use: "update",
		Run: cmd.UpdateProject,
	}, &cobra.Command{
		Use: "curd",
		Run: cmd.CreateCrudGroup,
	})
	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

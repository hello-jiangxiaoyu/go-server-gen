package main

import (
	"go-server-gen/cmd"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	newCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new project",
		Run:   cmd.NewProject,
	}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update idl config file",
		Run:   cmd.UpdateProject,
	}
	crudCmd := &cobra.Command{
		Use:   "curd",
		Short: "Create a new crud api",
		Run:   cmd.CreateCrudGroup,
	}
	cmd.InitCommand(rootCmd, newCmd, updateCmd, crudCmd)
	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

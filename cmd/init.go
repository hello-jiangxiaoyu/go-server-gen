package cmd

import "github.com/spf13/cobra"

var (
	LayoutPath = ""
	IdlPath    = ""

	CreateProjectName = ""
	CreateProjectTpl  = ""
	CrudServiceName   = ""
	CrudRouterPrefix  = ""
)

func InitCommand(rootCmd, newCmd, updateCmd, crudCmd *cobra.Command) {
	newCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	newCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")
	newCmd.PersistentFlags().StringVarP(&CreateProjectName, "module", "m", "", "go mod name")
	newCmd.PersistentFlags().StringVarP(&CreateProjectTpl, "tpl", "t", "login", "create template")

	updateCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	updateCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")

	crudCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	crudCmd.PersistentFlags().StringVarP(&CrudServiceName, "service", "s", "", "crud service name")
	crudCmd.PersistentFlags().StringVarP(&CrudRouterPrefix, "prefix", "p", "", "crud router prefix")

	rootCmd.AddCommand(newCmd, updateCmd, crudCmd)
}

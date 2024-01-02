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
	// 创建项目 new cmd
	newCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	newCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")
	newCmd.PersistentFlags().StringVarP(&CreateProjectName, "module", "m", "", "go mod name")
	newCmd.PersistentFlags().StringVarP(&CreateProjectTpl, "tpl", "t", "login", "create template")

	// 更新接口 update cmd
	updateCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	updateCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")

	// 添加crud
	crudCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	crudCmd.PersistentFlags().StringVarP(&CrudServiceName, "service", "s", "", "crud service name")
	crudCmd.PersistentFlags().StringVarP(&CrudRouterPrefix, "prefix", "p", "", "crud router prefix")

	rootCmd.AddCommand(newCmd, updateCmd, crudCmd)
}

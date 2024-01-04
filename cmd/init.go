package cmd

import "github.com/spf13/cobra"

var (
	LayoutPath = ""
	IdlPath    = ""
	ServerType = "gin"
	LogType    = "zap"

	CreateProjectName = "server"
	CrudServiceName   = ""
	CrudRouterPrefix  = ""
)

func InitCommand(rootCmd, newCmd, updateCmd, crudCmd *cobra.Command) {
	// 创建项目 new cmd
	newCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	newCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")
	newCmd.PersistentFlags().StringVarP(&CreateProjectName, "module", "m", "server", "go mod name")
	newCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "gin", "server type")
	newCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")

	// 更新接口 update cmd
	updateCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	updateCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")
	updateCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "gin", "server type")
	updateCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")

	// 添加crud
	crudCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	crudCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "gin", "server type")
	crudCmd.PersistentFlags().StringVar(&CrudRouterPrefix, "prefix", "", "crud router prefix")
	crudCmd.PersistentFlags().StringVar(&CrudServiceName, "service", "", "crud service name")
	crudCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")

	rootCmd.AddCommand(newCmd, updateCmd, crudCmd)
}

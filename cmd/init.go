package cmd

import "github.com/spf13/cobra"

var (
	LayoutPath = ""
	IdlPath    = ""
	ServerType = "gin"
	LogType    = "zap"
	OutputDir  = ""
	ForceWrite = false

	CreateProjectName = ""
	CrudServiceName   = ""
	CrudRouterPrefix  = ""
)

func InitCommand(rootCmd, newCmd, updateCmd, crudCmd *cobra.Command) {
	// 创建项目 new cmd
	newCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	newCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")
	newCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "gin", "server type")
	newCmd.PersistentFlags().StringVarP(&OutputDir, "output", "o", "", "output dir")
	newCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")
	newCmd.PersistentFlags().BoolVar(&ForceWrite, "force", false, "force write")

	// 更新接口 update cmd
	updateCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	updateCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")
	updateCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "gin", "server type")
	updateCmd.PersistentFlags().StringVarP(&OutputDir, "output", "o", "", "output dir")
	updateCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")

	// 添加crud
	crudCmd.PersistentFlags().StringVarP(&LayoutPath, "layout", "l", "", "layout path")
	crudCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "gin", "server type")
	crudCmd.PersistentFlags().StringVarP(&CrudRouterPrefix, "prefix", "p", "", "crud router prefix")
	crudCmd.PersistentFlags().StringVarP(&OutputDir, "output", "o", "", "output dir")
	crudCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")
	crudCmd.PersistentFlags().BoolVar(&ForceWrite, "force", false, "force write")

	rootCmd.AddCommand(newCmd, updateCmd, crudCmd)
}

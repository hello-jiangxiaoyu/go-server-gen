package cmd

import "github.com/spf13/cobra"

var (
	LayoutPath = ""
	IdlPath    = ""
	ServerType = ""
	LogType    = "zap"
	OutputDir  = ""
	ForceWrite = false
	WithTs     = false

	CreateProjectName = ""
	CrudServiceName   = ""
	RouterPrefix      = ""
	GormDSN           = ""
)

func InitCommand(rootCmd, newCmd, updateCmd, crudCmd, serverCmd *cobra.Command) {
	// 创建项目 new cmd
	newCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")
	newCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "gin", "server type")
	newCmd.PersistentFlags().StringVarP(&OutputDir, "output", "o", "", "output dir")
	newCmd.PersistentFlags().StringVar(&LayoutPath, "layout", "", "layout path")
	newCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")
	newCmd.PersistentFlags().BoolVar(&ForceWrite, "force", false, "force write")

	// 更新接口 update cmd
	updateCmd.PersistentFlags().StringVarP(&IdlPath, "idl", "i", "", "idl path")
	updateCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "", "server type")
	updateCmd.PersistentFlags().StringVarP(&OutputDir, "output", "o", "", "output dir")
	updateCmd.PersistentFlags().StringVar(&LayoutPath, "layout", "", "layout path")
	updateCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")
	updateCmd.PersistentFlags().BoolVar(&WithTs, "ts", false, "gen ts client code")

	// 添加crud
	crudCmd.PersistentFlags().StringVarP(&ServerType, "server", "s", "", "server type")
	crudCmd.PersistentFlags().StringVarP(&RouterPrefix, "prefix", "p", "", "crud router prefix")
	crudCmd.PersistentFlags().StringVarP(&OutputDir, "output", "o", "", "output dir")
	crudCmd.PersistentFlags().StringVar(&LayoutPath, "layout", "", "layout path")
	crudCmd.PersistentFlags().StringVar(&LogType, "log", "zap", "log type")
	crudCmd.PersistentFlags().BoolVar(&ForceWrite, "force", false, "force write")

	serverCmd.PersistentFlags().StringVar(&LayoutPath, "layout", "", "layout path")
	serverCmd.PersistentFlags().StringVar(&GormDSN, "dsn", "", "mysql dsn")

	rootCmd.AddCommand(newCmd, updateCmd, crudCmd, serverCmd)
}

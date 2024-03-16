package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-server-gen/cmd"
	"os"
	"runtime"
	"strings"
)

func main() {
	// panic handler
	defer func() {
		if err := recover(); err != nil {
			panicInfo := fmt.Sprintf("%v", err)
			for i := 3; i < 30; i++ {
				pc, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				if strings.Contains(file, "go-server-gen") {
					panicInfo += fmt.Sprintf("\n\t%s:%d %s", file, line, runtime.FuncForPC(pc).Name())
				}
			}
			fmt.Println("\npanic!!!", panicInfo)
		}
	}()

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
		Use:   "crud",
		Short: "Create a new crud api",
		Run:   cmd.CreateCrudGroup,
	}
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "web server",
		Run:   cmd.StartWebServer,
	}

	cmd.InitCommand(rootCmd, newCmd, updateCmd, crudCmd, serverCmd)
	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-server-gen/cmd"
	"os"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "gen",
		Run: func(cmd *cobra.Command, args []string) {
			os.Exit(1)
		},
	}
	rootCmd.AddCommand(&cobra.Command{
		Use: "create",
		Run: cmd.CreateProject,
	}, &cobra.Command{
		Use: "update",
		Run: cmd.UpdateProject,
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

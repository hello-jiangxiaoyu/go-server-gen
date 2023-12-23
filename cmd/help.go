package cmd

import "github.com/spf13/cobra"

func PrintHelp(_ *cobra.Command, _ []string) {
	println("gsg new")
	println("gsg update")
	println("gsg curd")
}

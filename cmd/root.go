package cmd

import (
	"github.com/spf13/cobra"
)

// TODO write cmd usages

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pure-live",
	Short: "pure-live",
	Long:  `pure-live`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

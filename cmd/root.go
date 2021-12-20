package cmd

import (
	"github.com/spf13/cobra"
)

// TODO write cmd usages

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pure-live",
	Short: "Make Live Pure Again",
	Long: `Make Live Pure Again.
No gift, fan group, pop-up window, only live, danmaku.
`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

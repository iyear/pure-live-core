package cmd

import (
	"github.com/fatih/color"
	"github.com/iyear/pure-live-core/global"
	"github.com/spf13/cobra"
)

// TODO write cmd usages

var version bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pure-live",
	Short: "Make Live Pure Again",
	Long: `Make Live Pure Again.
No gift, fan group, pop-up window, only live, danmaku.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			color.Blue("%s\n%s", global.Version, global.GetRuntime())
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "check the version of pure-live")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())

}

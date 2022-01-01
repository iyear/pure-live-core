package cmd

import (
	"github.com/iyear/pure-live-core/app/server"
	"github.com/spf13/cobra"
)

var (
	serverCfg  string
	accountCfg string
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the local server",
	Long:  `Start the local server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(serverCfg, accountCfg)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&serverCfg, "server", "s", "config/server.yaml", "server config file")
	runCmd.PersistentFlags().StringVarP(&accountCfg, "account", "a", "config/account.yaml", "account config file")
}

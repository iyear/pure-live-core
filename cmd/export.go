package cmd

import (
	"github.com/iyear/pure-live-core/app/export"
	"github.com/spf13/cobra"
)

var dbPath string
var savePath string

// exportCmd represents the run command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export Favorites And Favorites List Information",
	Long:  `Export Favorites And Favorites List Information`,
	Run: func(cmd *cobra.Command, args []string) {
		export.Export(dbPath, savePath)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.PersistentFlags().StringVarP(&dbPath, "db", "d", "data/data.db", "the path to data.db")
	exportCmd.PersistentFlags().StringVarP(&savePath, "path", "p", "export.xlsx", "the path to savePath")
}

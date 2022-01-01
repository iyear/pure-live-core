package cmd

import (
	"github.com/iyear/pure-live-core/app/get"
	"github.com/spf13/cobra"
)

var (
	plat    string
	room    string
	stream  string
	danmaku string
	roll    bool
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get live info",
	Long:  `Get live information, live stream, and danmaku stream`,
	Run: func(cmd *cobra.Command, args []string) {
		get.Get(plat, room, stream, danmaku, roll)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVarP(&plat, "plat", "p", "bilibili", "live platform")
	getCmd.PersistentFlags().StringVarP(&room, "room", "r", "6", "live room,it can be short string or real string")
	getCmd.PersistentFlags().StringVar(&stream, "stream", "", "download live stream to .flv file")
	getCmd.PersistentFlags().StringVar(&danmaku, "danmaku", "", "download live danmaku to .xlsx file")
	getCmd.PersistentFlags().BoolVar(&roll, "roll", false, "display danmaku content and scroll")
}

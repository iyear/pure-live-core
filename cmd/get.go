package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/service/srv_live"
	"github.com/spf13/cobra"
)

var (
	plat string
	room string
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get info",
	Long:  `get info`,
	Run: func(cmd *cobra.Command, args []string) {
		get()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVarP(&plat, "plat", "p", "bilibili", "live platform")
	getCmd.PersistentFlags().StringVarP(&room, "room", "r", "6", "live room")

}

func get() {
	info, err := srv_live.GetRoomInfo(plat, room)
	if err != nil {
		color.Red("[ERROR] can't get room info: %s", err)
		return
	}

	if info.Status == 0 {
		color.Yellow("[WARN] room is not online,so can't get the stream")
		infoOutput(info)
		return
	}
	url, err := srv_live.GetPlayURL(plat, info.Room)
	if err != nil {
		color.Red("[ERROR] can't get room info: %s", err)
		return
	}
	infoOutput(info)
	_, _ = fmt.Fprintf(color.Output, "Stream: %s", color.New(color.FgBlue).SprintFunc()(url.Origin))
}

func infoOutput(info *model.RoomInfo) {
	blue := color.New(color.FgBlue).SprintFunc()
	_, _ = fmt.Fprintf(color.Output, "Room: %s\nUpper: %s\nTitle: %s\nLink: %s\n",
		blue(info.Room),
		blue(info.Upper),
		blue(info.Title),
		blue(info.Link))
}

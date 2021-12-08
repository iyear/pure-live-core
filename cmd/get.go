package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/forwarder"
	"github.com/iyear/pure-live/service/srv_live"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	plat     string
	room     string
	download string
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
	getCmd.PersistentFlags().StringVarP(&download, "download", "d", "", "download live stream to .flv file")
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
	_, _ = fmt.Fprintf(color.Output, "Stream: %s\n", color.New(color.FgBlue).SprintFunc()(url.Origin))

	if download == "" {
		return
	}

	if err = dlStream(url.Type, url.Origin); err != nil {
		color.Red("[ERROR] can't download stream: %s", err)
		return
	}
	color.Blue("Download Live Stream Succ...")
}

func infoOutput(info *model.RoomInfo) {
	blue := color.New(color.FgBlue).SprintFunc()
	_, _ = fmt.Fprintf(color.Output, "Room: %s\nUpper: %s\nTitle: %s\nLink: %s\n",
		blue(info.Room),
		blue(info.Upper),
		blue(info.Title),
		blue(info.Link))
}

func dlStream(tp string, url string) error {
	writer := httpflv.FlvFileWriter{}

	if err := writer.Open(download); err != nil {
		return err
	}

	if err := writer.WriteFlvHeader(); err != nil {
		return err
	}

	err := forwarder.Pull(forwarder.GetIn(tp), url, func(tag httpflv.Tag) {
		_ = writer.WriteTag(tag)
	})
	if err != nil {
		return err
	}

	color.Yellow("[WARN] Ctrl + C to finish downloading")

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

	if err = writer.Dispose(); err != nil {
		return err
	}
	return nil

}

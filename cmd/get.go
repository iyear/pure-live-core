package cmd

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/client"
	"github.com/iyear/pure-live/pkg/conf"
	"github.com/iyear/pure-live/pkg/forwarder"
	"github.com/iyear/pure-live/pkg/util"
	"github.com/iyear/pure-live/service/srv_live"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/naza/pkg/nazalog"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"os"
	"os/signal"
	"sync"
	"time"
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
		get()
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

func get() {
	info, err := srv_live.GetRoomInfo(plat, room)
	if err != nil {
		color.Red("[ERROR] can't get room info: %s\n", err)
		return
	}

	if info.Status == 0 {
		color.Yellow("[WARN] room is not online,so can't get the stream\n")
		infoOutput(info)
		return
	}

	room = info.Room

	url, err := srv_live.GetPlayURL(plat, info.Room)
	if err != nil {
		color.Red("[ERROR] can't get room info: %s\n", err)
		return
	}
	infoOutput(info)
	_, _ = fmt.Fprintf(color.Output, "Stream: %s\n", color.New(color.FgBlue).SprintFunc()(url.Origin))

	color.Yellow("\n[WARN] Ctrl + C to finish downloading\n")
	ctx, stop := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	if stream != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := dlStream(ctx, url.Type, url.Origin); err != nil {
				color.Red("[ERROR] can't download stream: %s\n", err)
				return
			}
		}()
	}

	if danmaku != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := dlDanmaku(ctx); err != nil {
				color.Red("[ERROR] can't download danmaku: %s\n", err)
				return
			}
		}()
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	stop()
	wg.Wait()
}

func infoOutput(info *model.RoomInfo) {
	blue := color.New(color.FgBlue).SprintFunc()
	_, _ = fmt.Fprintf(color.Output, "Room: %s\nUpper: %s\nTitle: %s\nLink: %s\n",
		blue(info.Room),
		blue(info.Upper),
		blue(info.Title),
		blue(info.Link))
}

func dlStream(ctx context.Context, tp string, url string) error {
	if util.FileExists(stream) {
		color.Yellow("[WARN] file %s already exists, so skip stream downloading\n", stream)
		return nil
	}

	// 关闭nazalog的输出
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.IsToStdout = false
	})

	writer := httpflv.FlvFileWriter{}

	if err := writer.Open(stream); err != nil {
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

	<-ctx.Done()

	if err = writer.Dispose(); err != nil {
		return err
	}
	color.Blue("[INFO] Download Live Stream Succ...\n")
	return nil
}

func dlDanmaku(ctx context.Context) error {
	if util.FileExists(danmaku) {
		color.Yellow("[WARN] file %s already exists, so skip danmaku downloading\n", danmaku)
		return nil
	}

	cli, err := client.GetClient(plat)
	if err != nil {
		return err
	}
	defer cli.Stop()

	rev, err := srv_live.Serve(ctx, websocket.DefaultDialer, uuid.New().String(), cli, room)
	if err != nil {
		return err
	}

	file := excelize.NewFile()
	writer, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		return err
	}

	_ = writer.SetRow("A1", []interface{}{
		excelize.Cell{Value: "内容"},
		excelize.Cell{Value: "颜色(十进制)"},
		excelize.Cell{Value: "位置"},
		excelize.Cell{Value: "时间"},
	})

	count := 2
	for {
		select {
		case <-ctx.Done():
			if err = writer.Flush(); err != nil {
				return err
			}
			if err = file.SaveAs(danmaku); err != nil {
				return err
			}
			color.Blue("[INFO] Download Live Danmaku Succ...\n")
			return nil
		case transport := <-rev:
			if transport.Msg.Event() != conf.EventDanmaku {
				continue
			}
			dm := transport.Msg.(*model.MsgDanmaku)
			if roll {
				color.Cyan("%s %d\n", dm.Content, dm.Color)
			}
			cell, err := excelize.CoordinatesToCellName(1, count)
			if err != nil {
				return err
			}
			if err = writer.SetRow(cell, []interface{}{
				excelize.Cell{Value: dm.Content},
				excelize.Cell{Value: dm.Color},
				excelize.Cell{Value: util.DmMode2Desc(dm.Type)},
				excelize.Cell{Value: time.Now().Format("2006-01-02 15:04:05")},
			}); err != nil {
				return err
			}
			count++
		}
	}
}

package get

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/forwarder"
	"github.com/iyear/pure-live-core/pkg/util"
	"github.com/iyear/pure-live-core/service/svc_live"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/naza/pkg/nazalog"
	"github.com/xuri/excelize/v2"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Get(plat, room, stream, danmaku string, roll bool) {
	info, err := svc_live.GetRoomInfo(plat, room)
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

	url, err := svc_live.GetPlayURL(plat, info.Room)
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
			if err := dlStream(ctx, url.Type, url.Origin, stream); err != nil {
				color.Red("[ERROR] can't download stream: %s\n", err)
				return
			}
		}()
	}

	if danmaku != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := dlDanmaku(ctx, plat, room, danmaku, roll); err != nil {
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

func dlStream(ctx context.Context, tp string, url string, path string) error {
	if util.FileExists(path) {
		color.Yellow("[WARN] file %s already exists, so skip stream downloading\n", path)
		return nil
	}

	// 关闭nazalog的输出
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.IsToStdout = false
	})

	writer := httpflv.FlvFileWriter{}

	if err := writer.Open(path); err != nil {
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

func dlDanmaku(ctx context.Context, plat, room, path string, roll bool) error {
	if util.FileExists(path) {
		color.Yellow("[WARN] file %s already exists, so skip danmaku downloading\n", path)
		return nil
	}

	cli, err := client.GetClient(plat)
	if err != nil {
		return err
	}
	defer cli.Stop()

	rev, err := svc_live.Serve(ctx, websocket.DefaultDialer, uuid.New().String(), cli, room)
	if err != nil {
		return err
	}

	file := excelize.NewFile()
	file.NewSheet("弹幕")
	file.DeleteSheet("Sheet1")
	writer, err := file.NewStreamWriter("弹幕")
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
			if err = file.SaveAs(path); err != nil {
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

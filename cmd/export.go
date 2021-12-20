package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/db"
	"github.com/iyear/pure-live/pkg/util"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"time"
)

var dbFile string
var saveFile string

// exportCmd represents the run command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Start the local server",
	Long:  `Start the local server`,
	Run: func(cmd *cobra.Command, args []string) {
		export()
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.PersistentFlags().StringVarP(&dbFile, "db", "d", "data/data.db", "the path to data.db")
	exportCmd.PersistentFlags().StringVarP(&saveFile, "path", "p", "export.xlsx", "the path to saveFile")
}

func export() {
	if !util.FileExists(dbFile) {
		color.Red("[ERROR] %s is not exists", dbFile)
		return
	}

	if util.FileExists(saveFile) {
		color.Yellow("[WARN] file %s already exists, so skip exporting", saveFile)
		return
	}

	if err := db.Init(dbFile); err != nil {
		color.Red("[ERROR] failed to init db.ERROR: %s", err)
		return
	}

	var lists []*model.FavoritesList

	if err := db.SQLite.Find(&lists).Error; err != nil {
		color.Red("[ERROR] failed to get fav lists.ERROR: %s", err)
		return
	}

	file := excelize.NewFile()
	for _, list := range lists {
		file.NewSheet(list.Title)
		writer, err := file.NewStreamWriter(list.Title)
		if err != nil {
			color.Red("[ERROR] failed to create new excel writer.ERROR: %s", err)
			return
		}

		_ = writer.SetRow("A1", []interface{}{
			excelize.Cell{Value: fmt.Sprintf("ID: %d", list.ID)},
			excelize.Cell{Value: fmt.Sprintf("标题: %s", list.Title)},
			excelize.Cell{Value: fmt.Sprintf("排序: %d", list.Order)},
			excelize.Cell{Value: fmt.Sprintf("创建时间: %s", time.Unix(list.CreatedAt, 0).Format("2006-01-02 15:04:05"))},
			excelize.Cell{Value: fmt.Sprintf("最后更新: %s", time.Unix(list.UpdatedAt, 0).Format("2006-01-02 15:04:05"))},
		})

		_ = writer.SetRow("A2", []interface{}{})

		_ = writer.SetRow("A3", []interface{}{
			excelize.Cell{Value: "ID"},
			excelize.Cell{Value: "平台"},
			excelize.Cell{Value: "房间号"},
			excelize.Cell{Value: "主播"},
			excelize.Cell{Value: "排序"},
			excelize.Cell{Value: "创建时间"},
			excelize.Cell{Value: "最后更新"},
		})

		var favs []*model.Favorite
		if err = db.SQLite.Where("fid = ?", list.ID).Find(&favs).Error; err != nil {
			color.Red("[ERROR] failed to get favs.ERROR: %s", err)
			return
		}

		if err = writeFavs(writer, favs); err != nil {
			color.Red("[ERROR] failed to write favs.ERROR: %s", err)
			return
		}

		if err = writer.Flush(); err != nil {
			color.Red("[ERROR] failed to flush excel writer.ERROR: %s", err)
			return
		}
	}

	file.DeleteSheet("Sheet1")
	if err := file.SaveAs(saveFile); err != nil {
		color.Red("[ERROR] failed to save file.ERROR: %s", err)
		return
	}

}

func writeFavs(writer *excelize.StreamWriter, favs []*model.Favorite) error {
	var err error
	count := 4 // 前面已经有3行了
	for _, fav := range favs {
		cell := ""
		if cell, err = excelize.CoordinatesToCellName(1, count); err != nil {
			return err
		}
		if err = writer.SetRow(cell, []interface{}{
			excelize.Cell{Value: fav.ID},
			excelize.Cell{Value: util.Plat2Desc(fav.Plat)},
			excelize.Cell{Value: fav.Room},
			excelize.Cell{Value: fav.Upper},
			excelize.Cell{Value: fav.Order},
			excelize.Cell{Value: time.Unix(fav.CreatedAt, 0).Format("2006-01-02 15:04:05")},
			excelize.Cell{Value: time.Unix(fav.UpdatedAt, 0).Format("2006-01-02 15:04:05")},
		}); err != nil {
			return err
		}
		count++
	}
	return nil
}

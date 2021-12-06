package db

import (
	"github.com/iyear/pure-live/conf"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/sqlite"
	"gorm.io/gorm"
	"path"
	"time"
)

var SQLite *gorm.DB

func Init() error {
	var err error
	SQLite, err = gorm.Open(sqlite.Open(path.Join(conf.C.Server.Path, "data.db")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return err
	}

	if err = SQLite.AutoMigrate(&model.FavoritesList{}, &model.Favorite{}); err != nil {
		return err
	}

	// 创建默认收藏夹
	if err = SQLite.FirstOrCreate(&model.FavoritesList{}, &model.FavoritesList{
		ID:    1,
		Title: "默认收藏夹",
		Order: 0,
	}).Error; err != nil {
		return err
	}
	return nil
}

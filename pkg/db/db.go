package db

import (
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/sqlite"
	"gorm.io/gorm"
	"time"
)

func Init(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&model.FavoritesList{}, &model.Favorite{}); err != nil {
		return nil, err
	}

	// 创建默认收藏夹
	if err = db.FirstOrCreate(&model.FavoritesList{}, &model.FavoritesList{
		ID:    1,
		Title: "默认收藏夹",
		Order: 0,
	}).Error; err != nil {
		return nil, err
	}
	return db, nil
}

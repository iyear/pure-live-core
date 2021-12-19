package srv_fav

import (
	"fmt"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/db"
)

func AddFavList(title string, order int) (*model.FavoritesList, error) {
	result := model.FavoritesList{
		Title: title,
		Order: order,
	}
	if err := db.SQLite.Create(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func GetAllFavLists() ([]*model.FavoritesList, error) {
	var result []*model.FavoritesList
	if err := db.SQLite.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func DelFavList(id uint64) error {
	// 不允许删除默认收藏夹
	if id == 1 {
		return fmt.Errorf("default fav list cannot be deleted")
	}
	if err := db.SQLite.First(&model.FavoritesList{ID: id}).Limit(1).Error; err != nil {
		return err
	}
	if err := db.SQLite.Delete(&model.FavoritesList{ID: id}).Error; err != nil {
		return err
	}
	return nil
}

func EditFavList(id uint64, title string, order int) (*model.FavoritesList, error) {
	r := model.FavoritesList{ID: id}
	if err := db.SQLite.First(&r).Limit(1).Error; err != nil {
		return nil, err
	}
	if err := db.SQLite.Model(&r).Updates(map[string]interface{}{"title": title, "order": order}).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func GetFavList(id uint64) (*model.FavoritesList, []*model.Favorite, error) {
	var (
		list = model.FavoritesList{}
		favs []*model.Favorite
	)
	if err := db.SQLite.First(&list, id).Limit(1).Error; err != nil {
		return nil, nil, err
	}
	if err := db.SQLite.Where("fid = ?", id).Find(&favs).Error; err != nil {
		return nil, nil, err
	}
	return &list, favs, nil
}

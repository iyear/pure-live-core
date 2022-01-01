package svc_fav

import (
	"github.com/iyear/pure-live-core/global"
	"github.com/iyear/pure-live-core/model"
)

func GetFav(id uint64) (*model.Favorite, error) {
	fav := model.Favorite{}
	if err := global.DB.First(&fav, id).Limit(1).Error; err != nil {
		return nil, err
	}
	return &fav, nil
}
func AddFav(fid uint64, order int, plat string, room string, upper string) (*model.Favorite, error) {
	fav := model.Favorite{
		FID:   fid,
		Order: order,
		Plat:  plat,
		Room:  room,
		Upper: upper,
	}
	if err := global.DB.Create(&fav).Error; err != nil {
		return nil, err
	}
	return &fav, nil
}
func DelFav(id uint64) error {
	if err := global.DB.First(&model.Favorite{ID: id}).Limit(1).Error; err != nil {
		return err
	}
	if err := global.DB.Where("id = ?", id).Delete(&model.Favorite{}).Error; err != nil {
		return err
	}
	return nil
}

func EditFav(id uint64, order int, plat string, room string, upper string) (*model.Favorite, error) {
	r := model.Favorite{ID: id}
	if err := global.DB.First(&r).Limit(1).Error; err != nil {
		return nil, err
	}
	if err := global.DB.Model(&r).Updates(map[string]interface{}{"order": order, "plat": plat, "room": room, "upper": upper}).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

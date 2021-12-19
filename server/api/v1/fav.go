package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/e"
	"github.com/iyear/pure-live/server/api"
	"github.com/iyear/pure-live/service/srv_fav"
	"go.uber.org/zap"
)

func AddFavList(c *gin.Context) {
	req := struct {
		Title string `form:"title" binding:"required,min=2,max=60" json:"title"`
		Order int    `form:"order" binding:"required,gte=0,lte=100" json:"order"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}
	result, err := srv_fav.AddFavList(req.Title, req.Order)
	if err != nil {
		api.RespFmt(c, e.ErrorAddFavList, err, nil)
		zap.S().Warnw("failed to add fav list", "error", err, "req", req)
		return
	}
	api.RespFmt(c, e.Success, nil, result)
}
func GetAllFavLists(c *gin.Context) {
	result, err := srv_fav.GetAllFavLists()
	if err != nil {
		api.RespFmt(c, e.ErrorGetAllFavList, err, nil)
		zap.S().Warnw("failed to get all fav lists", "error", err)
		return
	}
	api.RespFmt(c, e.Success, nil, result)
}
func DelFavList(c *gin.Context) {
	req := struct {
		ID uint64 `form:"id" binding:"required" json:"id"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}
	if err := srv_fav.DelFavList(req.ID); err != nil {
		api.RespFmt(c, e.ErrorDelFavList, err, nil)
		zap.S().Warnw("failed to del fav list", "error", err, "req", req)
		return
	}
	api.RespFmt(c, e.Success, nil, nil)
}
func EditFavList(c *gin.Context) {
	req := struct {
		ID    uint64 `form:"id" binding:"required" json:"id"`
		Title string `form:"title" binding:"required,min=2,max=40" json:"title"`
		Order int    `form:"order" binding:"required,gte=0,lte=100" json:"order"`
	}{}

	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}
	r, err := srv_fav.EditFavList(req.ID, req.Title, req.Order)
	if err != nil {
		api.RespFmt(c, e.ErrorEditFavList, err, nil)
		zap.S().Warnw("failed to edit fav list", "error", err, "req", req)
		return
	}
	api.RespFmt(c, e.Success, nil, r)
}
func GetFavList(c *gin.Context) {
	req := struct {
		ID uint64 `form:"id" binding:"required"`
	}{}

	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}
	list, favs, err := srv_fav.GetFavList(req.ID)
	if err != nil {
		api.RespFmt(c, e.ErrorGetFavList, err, nil)
		zap.S().Warnw("failed to get fav list", "error", err, "req", req)
		return
	}
	api.RespFmt(c, e.Success, nil, &struct {
		*model.FavoritesList
		Favorites []*model.Favorite `json:"favorites"`
	}{FavoritesList: list, Favorites: favs})
}
func GetFav(c *gin.Context) {
	req := struct {
		ID uint64 `form:"id" binding:"required" json:"id"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}

	r, err := srv_fav.GetFav(req.ID)
	if err != nil {
		api.RespFmt(c, e.ErrorGetFav, err, nil)
		zap.S().Warnw("failed to get fav", "error", err)
		return
	}
	api.RespFmt(c, e.Success, nil, r)
}
func AddFav(c *gin.Context) {
	req := struct {
		FID   uint64 `form:"fid" binding:"required" json:"fid"`
		Order int    `form:"order" binding:"required,gte=0,lte=100" json:"order"`
		Plat  string `form:"plat" binding:"required,max=15" json:"plat"`
		Room  string `form:"room" binding:"required" json:"room"`
		Upper string `form:"upper" binding:"required" json:"upper"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}
	result, err := srv_fav.AddFav(req.FID, req.Order, req.Plat, req.Room, req.Upper)
	if err != nil {
		api.RespFmt(c, e.ErrorAddFav, err, nil)
		zap.S().Warnw("failed to add fav", "error", err, "req", req)
		return
	}
	api.RespFmt(c, e.Success, nil, result)
}
func DelFav(c *gin.Context) {
	req := struct {
		ID uint64 `form:"id" binding:"required" json:"id"`
	}{}

	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}

	if err := srv_fav.DelFav(req.ID); err != nil {
		api.RespFmt(c, e.ErrorDelFav, err, nil)
		zap.S().Warnw("failed to del fav", "error", err)
		return
	}
	api.RespFmt(c, e.Success, nil, nil)
}
func EditFav(c *gin.Context) {
	req := struct {
		ID    uint64 `form:"id" binding:"required" json:"id"`
		Order int    `form:"order" binding:"required,gte=0,lte=100" json:"order"`
		Plat  string `form:"plat" binding:"required,max=15" json:"plat"`
		Room  string `form:"room" binding:"required,gte=0,lte=100" json:"room"`
		Upper string `form:"upper" binding:"required,gte=0,lte=100" json:"upper"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}

	r, err := srv_fav.EditFav(req.ID, req.Order, req.Plat, req.Room, req.Upper)
	if err != nil {
		api.RespFmt(c, e.ErrorEditFav, err, nil)
		zap.S().Warnw("failed to edit fav", "error", err)
		return
	}
	api.RespFmt(c, e.Success, nil, r)
}

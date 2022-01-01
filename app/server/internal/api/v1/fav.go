package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/ecode"
	"github.com/iyear/pure-live-core/pkg/format"
	"github.com/iyear/pure-live-core/service/svc_fav"
	"go.uber.org/zap"
)

func AddFavList(c *gin.Context) {
	req := struct {
		Title string `form:"title" binding:"required,min=2,max=60" json:"title"`
		Order int    `form:"order" binding:"required,gte=0,lte=100" json:"order"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}
	result, err := svc_fav.AddFavList(req.Title, req.Order)
	if err != nil {
		format.HTTP(c, ecode.ErrorAddFavList, err, nil)
		zap.S().Warnw("failed to add fav list", "error", err, "req", req)
		return
	}
	format.HTTP(c, ecode.Success, nil, result)
}
func GetAllFavLists(c *gin.Context) {
	result, err := svc_fav.GetAllFavLists()
	if err != nil {
		format.HTTP(c, ecode.ErrorGetAllFavList, err, nil)
		zap.S().Warnw("failed to get all fav lists", "error", err)
		return
	}
	format.HTTP(c, ecode.Success, nil, result)
}
func DelFavList(c *gin.Context) {
	req := struct {
		ID uint64 `form:"id" binding:"required" json:"id"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}
	if err := svc_fav.DelFavList(req.ID); err != nil {
		format.HTTP(c, ecode.ErrorDelFavList, err, nil)
		zap.S().Warnw("failed to del fav list", "error", err, "req", req)
		return
	}
	format.HTTP(c, ecode.Success, nil, nil)
}
func EditFavList(c *gin.Context) {
	req := struct {
		ID    uint64 `form:"id" binding:"required" json:"id"`
		Title string `form:"title" binding:"required,min=2,max=40" json:"title"`
		Order int    `form:"order" binding:"required,gte=0,lte=100" json:"order"`
	}{}

	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}
	r, err := svc_fav.EditFavList(req.ID, req.Title, req.Order)
	if err != nil {
		format.HTTP(c, ecode.ErrorEditFavList, err, nil)
		zap.S().Warnw("failed to edit fav list", "error", err, "req", req)
		return
	}
	format.HTTP(c, ecode.Success, nil, r)
}
func GetFavList(c *gin.Context) {
	req := struct {
		ID uint64 `form:"id" binding:"required"`
	}{}

	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}
	list, favs, err := svc_fav.GetFavList(req.ID)
	if err != nil {
		format.HTTP(c, ecode.ErrorGetFavList, err, nil)
		zap.S().Warnw("failed to get fav list", "error", err, "req", req)
		return
	}
	format.HTTP(c, ecode.Success, nil, &struct {
		*model.FavoritesList
		Favorites []*model.Favorite `json:"favorites"`
	}{FavoritesList: list, Favorites: favs})
}
func GetFav(c *gin.Context) {
	req := struct {
		ID uint64 `form:"id" binding:"required" json:"id"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}

	r, err := svc_fav.GetFav(req.ID)
	if err != nil {
		format.HTTP(c, ecode.ErrorGetFav, err, nil)
		zap.S().Warnw("failed to get fav", "error", err)
		return
	}
	format.HTTP(c, ecode.Success, nil, r)
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
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}
	result, err := svc_fav.AddFav(req.FID, req.Order, req.Plat, req.Room, req.Upper)
	if err != nil {
		format.HTTP(c, ecode.ErrorAddFav, err, nil)
		zap.S().Warnw("failed to add fav", "error", err, "req", req)
		return
	}
	format.HTTP(c, ecode.Success, nil, result)
}
func DelFav(c *gin.Context) {
	req := struct {
		ID uint64 `form:"id" binding:"required" json:"id"`
	}{}

	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}

	if err := svc_fav.DelFav(req.ID); err != nil {
		format.HTTP(c, ecode.ErrorDelFav, err, nil)
		zap.S().Warnw("failed to del fav", "error", err)
		return
	}
	format.HTTP(c, ecode.Success, nil, nil)
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
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}

	r, err := svc_fav.EditFav(req.ID, req.Order, req.Plat, req.Room, req.Upper)
	if err != nil {
		format.HTTP(c, ecode.ErrorEditFav, err, nil)
		zap.S().Warnw("failed to edit fav", "error", err)
		return
	}
	format.HTTP(c, ecode.Success, nil, r)
}

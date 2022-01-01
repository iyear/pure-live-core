package model

// FavoritesList fav list
type FavoritesList struct {
	ID    uint64 `gorm:"primaryKey;column:id;index" json:"id"`
	Title string `gorm:"not null;unique;column:title" json:"title"`
	Order int    `gorm:"not null;column:order" json:"order"`
	TimeHook
}

// Favorite fav for live room
type Favorite struct {
	ID    uint64 `gorm:"primaryKey;column:id;index" json:"id"` // ID
	FID   uint64 `gorm:"not null;column:fid;index" json:"fid"` // 收藏夹ID
	Order int    `gorm:"not null;column:order" json:"order"`   // 排序
	Plat  string `gorm:"not null;column:plat" json:"plat"`     // 平台
	Room  string `gorm:"not null;column:room" json:"room"`     // 房间名
	Upper string `gorm:"not null;column:upper" json:"upper"`   // 主播名
	TimeHook
}

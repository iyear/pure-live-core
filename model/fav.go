package model

type FavoritesList struct {
	ID    uint64 `gorm:"primaryKey;column:id" json:"id"`
	Title string `gorm:"not null;unique;column:title" json:"title"`
	Order int    `gorm:"not null;column:order" json:"order"`
	TimeHook
}

type Favorite struct {
	ID    uint64 `gorm:"primaryKey;column:id" json:"id"`     // ID
	FID   uint64 `gorm:"not null;column:fid" json:"fid"`     // 收藏夹ID
	Order int    `gorm:"not null;column:order" json:"order"` // 排序
	Plat  string `gorm:"not null;column:plat" json:"plat"`   // 平台
	Room  string `gorm:"not null;column:room" json:"room"`   // 房间名
	Upper string `gorm:"not null;column:upper" json:"upper"` // 主播名
	TimeHook
}

package model

type TimeHook struct {
	CreatedAt int64 `gorm:"not null;column:created_at" json:"created_at"`
	UpdatedAt int64 `gorm:"not null;column:updated_at" json:"updated_at"`
}

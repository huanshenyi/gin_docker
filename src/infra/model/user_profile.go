package model

import "time"

type UserProfile struct {
	ID         int       `gorm:"column:id"`
	Email      string    `gorm:"column:email"`
	Sex        int       `gorm:"column:sex"`
	LivingArea string    `gorm:"column:living_area"` // 生活エリア
	Age        int       `gorm:"column:age"`         // 年齢
	Appeal     string    `gorm:"column:appeal"`      // 一言アピール
	Profession string    `gorm:"column:profession"`  // 職業
	CreatedAt  time.Time `gorm:"column:created"`
	UpdatedAt  time.Time `gorm:"column:modified"`
	UserID     int       `gorm:"column:user_id"`
}

func (u UserProfile) TableName() string {
	return "user_profiles"
}

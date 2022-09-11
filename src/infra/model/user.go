package model

import "time"

type User struct {
	ID           int           `gorm:"column:id"`
	UserName     string        `gorm:"column:username"`
	Icon         string        `gorm:"icon"`
	CreatedAt    time.Time     `gorm:"column:created"`  // created
	UpdatedAt    time.Time     `gorm:"column:modified"` // modified
	UserAuths    []UserAuth    `gorm:"foreignKey:UserID"`
	Recruitments []Recruitment `gorm:"foreignKey:UserID"`
	UserProfile  UserProfile
}

func (m *User) TableName() string {
	return "users"
}

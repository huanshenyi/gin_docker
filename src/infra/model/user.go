package model

import "time"

type User struct {
	ID           int           `gorm:"column:id"`
	UserName     string        `gorm:"column:username"`
	Icon         string        `gorm:"icon"`
	Group        int           `gorm:"group;comment:'1 admin 2 user...'"`
	CreatedAt    time.Time     `gorm:"column:created"`  // created
	UpdatedAt    time.Time     `gorm:"column:modified"` // modified
	UserAuths    []UserAuth    `gorm:"foreignKey:UserID"`
	Recruitments []Recruitment `gorm:"foreignKey:UserID"`
	UserProfile  UserProfile
}

func (m *User) TableName() string {
	return "users"
}

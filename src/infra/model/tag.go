package model

import "time"

type Tag struct {
	ID        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Status    int       `gorm:"column:status"`   // 1, 0
	CreatedAt time.Time `gorm:"column:created"`  // created
	UpdatedAt time.Time `gorm:"column:modified"` // modified
}

func (m *Tag) TableName() string {
	return "tags"
}

package model

import (
	"gin_docker/src/domain/tag"
	"time"
)

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

func (m *Tag) ToDomain() tag.TagData {
	return tag.TagData{
		ID:     m.ID,
		Name:   m.Name,
		Status: m.Status,
	}
}

package model

import (
	"time"

	"gin_docker/src/domain"
)

type Tag struct {
	ID           int           `gorm:"column:id"`
	Name         string        `gorm:"column:name"`
	Status       int           `gorm:"column:status"`   // 1, 0
	CreatedAt    time.Time     `gorm:"column:created"`  // created
	UpdatedAt    time.Time     `gorm:"column:modified"` // modified
	Recruitments []Recruitment `gorm:"many2many:recruitment_tags"`
}

func (m *Tag) TableName() string {
	return "tags"
}

func (m *Tag) ToDomain() domain.TagData {
	return domain.TagData{
		ID:     m.ID,
		Name:   m.Name,
		Status: m.Status,
	}
}

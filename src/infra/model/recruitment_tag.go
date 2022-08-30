package model

import "time"

// RecruitmentTag - 募集に紐づくタグ情報を保存
type RecruitmentTag struct {
	ID            int       `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	RecruitmentID int       `gorm:"column:recruitment_id"`
	TagID         int       `gorm:"column:tag_id"`
	UpdatedAt     time.Time `gorm:"column:modified"`
	CreatedAt     time.Time `gorm:"column:created"`
}

// TableName sets the original table name
func (t *RecruitmentTag) TableName() string {
	return "recruitment_tags"
}

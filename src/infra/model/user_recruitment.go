package model

import "time"

// UserRecruitment 募集に応募する場合のユーザーと募集の関係
type UserRecruitment struct {
	ID            int       `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	UserID        int       `gorm:"column:user_id"`
	RecruitmentID int       `gorm:"column:recruitment_id"`
	UpdatedAt     time.Time `gorm:"column:modified"`
	CreatedAt     time.Time `gorm:"column:created"`
}

func (m *UserRecruitment) TableName() string {
	return "user_recruitments"
}

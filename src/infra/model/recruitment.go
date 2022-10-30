package model

import (
	"time"

	"gin_docker/src/domain"
)

type Recruitment struct {
	ID          int       `gorm:"column:id;primary_key"`
	Title       string    `gorm:"column:title"`
	Place       string    `gorm:"column:place"`       // 場所
	Start       time.Time `gorm:"column:start"`       // 開始時間
	End         time.Time `gorm:"column:end"`         // 終了時間
	Content     string    `gorm:"column:content"`     // 募集説明
	Paid        bool      `gorm:"column:paid"`        // 有償か
	Reward      string    `gorm:"column:reward"`      // 報酬内容
	MemberLimit int       `gorm:"column:memberLimit"` // 募集人数
	Type        string    `gorm:"column:type"`        // 何タイプの募集
	UserID      int       `gorm:"column:user_id"`     // オーナー
	CreatedAt   time.Time `gorm:"column:created"`
	UpdatedAt   time.Time `gorm:"column:modified"`
	Tags        []Tag     `gorm:"many2many:recruitment_tags;constraint:OnDelete:CASCADE"`
}

func (m *Recruitment) TableName() string {
	return "recruitments"
}

func (m *Recruitment) ToDomain() domain.Recruitment {
	tags := make([]domain.TagData, len(m.Tags))
	for k, i := range m.Tags {
		tags[k] = i.ToDomain()
	}
	return domain.Recruitment{
		ID:          m.ID,
		Title:       m.Title,
		Place:       m.Place,
		Start:       m.Start,
		End:         m.End,
		Content:     m.Content,
		Paid:        m.Paid,
		Reward:      m.Reward,
		MemberLimit: m.MemberLimit,
		Type:        domain.RecruitmentType(m.Type),
		UserID:      m.UserID,
		Tags:        tags,
	}
}

func RecruitmentTypeToSQL(value domain.RecruitmentType) string {
	switch value {
	case domain.RecruitmentTypeDefault:
		return "recruitment"
	case domain.RecruitmentTypeFreeTime:
		return "freeTime"
	default:
		return "recruitment"
	}
}

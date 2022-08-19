package recruitment

import (
	"time"

	"gin_docker/src/domain"
)

// CreateInput 募集追加に必要な要素
type CreateInput struct {
	UserID      int
	Title       string                 `json:"title" validate:"required"`
	Place       string                 `json:"place" validate:"required"`
	Start       time.Time              `json:"start" validate:"required"`
	End         time.Time              `json:"end" validate:"required"`
	Content     string                 `json:"content" validate:"required"`
	Paid        bool                   `json:"paid" validate:"required"`
	Reward      string                 `json:"reward" validate:"required"`
	MemberLimit int                    `json:"memberLimit" validate:"required"`
	Type        domain.RecruitmentType `json:"type" default:"recruitment"`
}

func (i *interactor) Create(input CreateInput) error {
	err := i.repository.CreateRecruitment(i.tx, domain.Recruitment{
		Title:       input.Title,
		Place:       input.Place,
		Start:       input.Start,
		End:         input.End,
		Content:     input.Content,
		Paid:        input.Paid,
		Reward:      input.Reward,
		MemberLimit: input.MemberLimit,
		UserID:      input.UserID,
		Type:        input.Type,
	})
	if err != nil {
		return err
	}
	return nil
}

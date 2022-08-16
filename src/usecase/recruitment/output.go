package recruitment

import (
	"gin_docker/src/domain"
	"gin_docker/src/domain/recruitment"
	"time"
)

type ListOutput struct {
	Recruitments []Recruitment `json:"recruitments"`
}

type Recruitment struct {
	ID          int                    `json:"id"`
	Title       string                 `json:"title"`
	Place       string                 `json:"place"`
	Start       time.Time              `json:"start"`
	End         time.Time              `json:"end"`
	Content     string                 `json:"content"`
	Paid        bool                   `json:"paid"`
	Reward      string                 `json:"reward"`
	MemberLimit int                    `json:"memberLimit"`
	Type        domain.RecruitmentType `json:"type"`
}

func ConvertRecruitmentOutput(recruitments recruitment.Recruitments) (out ListOutput) {
	recruitmentList := make([]Recruitment, len(recruitments.Recruitments))

	for i, r := range recruitments.Recruitments {
		recruitmentList[i] = Recruitment{
			ID:          r.ID,
			Title:       r.Title,
			Place:       r.Place,
			Start:       r.Start,
			End:         r.End,
			Content:     r.Content,
			Paid:        r.Paid,
			Reward:      r.Reward,
			MemberLimit: r.MemberLimit,
			Type:        r.Type,
		}
	}
	out = ListOutput{
		Recruitments: recruitmentList,
	}
	return
}

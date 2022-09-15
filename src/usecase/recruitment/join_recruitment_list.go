package recruitment

import (
	"time"

	"gin_docker/src/domain"
)

type JoinListInput struct {
	UserID int
	Limit  int `form:"limit,default=10"`
	Page   int `form:"page,default=1"`
}

type JoinListOutput struct {
	JoinListRecruitments []JoinListRecruitment `json:"joinListRecruitments"`
	TotalPage            int                   `json:"totalPage"`
	TotalCount           int                   `json:"totalCount"`
}

type JoinListRecruitment struct {
	Recruitment RecruitmentInfo  `json:"recruitment"`
	Owner       RecruitmentOwner `json:"owner"`
}

type RecruitmentInfo struct {
	ID          int                    `json:"id"`
	Title       string                 `json:"title"`
	Place       string                 `json:"place"`
	Start       time.Time              `json:"start"`
	End         time.Time              `json:"end"`
	Content     string                 `json:"content"`
	Paid        bool                   `json:"paid"`
	Reward      string                 `json:"reward"`
	MemberLimit int                    `json:"member_limit"`
	UserID      int                    `json:"user_id"`
	Type        domain.RecruitmentType `json:"type"`
	Tags        []RecruitmentTag       `json:"tags"`
}

type RecruitmentOwner struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}

type RecruitmentTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (i *interactor) JoinList(input JoinListInput) (output JoinListOutput, err error) {
	res, err := i.repository.JoinListRecruitment(i.tx, input.UserID, input.Page, input.Limit)
	if err != nil {
		return JoinListOutput{}, err
	}
	jList := make([]JoinListRecruitment, len(res.Recruitment))
	for i, k := range res.Recruitment {
		tags := make([]RecruitmentTag, len(k.Recruitment.Tags))
		for i, t := range k.Recruitment.Tags {
			tags[i] = RecruitmentTag{
				ID:   t.ID,
				Name: t.Name,
			}
		}
		jList[i] = JoinListRecruitment{
			Recruitment: RecruitmentInfo{
				ID:          k.Recruitment.ID,
				Title:       k.Recruitment.Title,
				Place:       k.Recruitment.Place,
				Start:       k.Recruitment.Start,
				End:         k.Recruitment.End,
				Content:     k.Recruitment.Content,
				Paid:        k.Recruitment.Paid,
				Reward:      k.Recruitment.Reward,
				MemberLimit: k.Recruitment.MemberLimit,
				UserID:      k.Recruitment.UserID,
				Type:        k.Recruitment.Type,
				Tags:        tags,
			},
			Owner: RecruitmentOwner{
				ID:       k.Owner.ID,
				UserName: k.Owner.UserName,
				Icon:     k.Owner.Icon,
			},
		}
	}
	output.JoinListRecruitments = jList
	output.TotalCount = res.TotalCount
	output.TotalPage = res.TotalPage

	return output, nil
}

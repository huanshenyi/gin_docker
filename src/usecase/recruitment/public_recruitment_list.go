package recruitment

import "gin_docker/src/domain"

type PublicListInput struct {
	Tag   string `form:"tag"`
	Type  string `form:"type,default=recruitment" validate:"oneof= recruitment freetime"`
	Limit int    `form:"limit,default=10"`
	Page  int    `form:"page,default=1"`
}

type PublicListOutput struct {
	Recruitments []JoinListRecruitment `json:"recruitments"`
	TotalPage    int                   `json:"totalPage"`
	TotalCount   int                   `json:"totalCount"`
}

func (i *interactor) PublicList(input PublicListInput) (output PublicListOutput, err error) {
	res, err := i.repository.PublicList(i.tx, domain.ToRecruitmentType(input.Type), input.Tag, input.Page, input.Limit)
	if err != nil {
		return PublicListOutput{}, err
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
				Created:     k.Recruitment.Created,
				Tags:        tags,
			},
			Owner: RecruitmentOwner{
				ID:       k.Owner.ID,
				UserName: k.Owner.UserName,
				Icon:     k.Owner.Icon,
			},
		}
	}
	output.Recruitments = jList
	output.TotalCount = res.TotalCount
	output.TotalPage = res.TotalPage

	return output, nil
}

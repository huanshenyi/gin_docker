package recruitment

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
	return PublicListOutput{}, nil
}

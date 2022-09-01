package domain

import "time"

type RecruitmentType string

const (
	RecruitmentTypeDefault  RecruitmentType = "recruitment"
	RecruitmentTypeFreeTime RecruitmentType = "freeTime"
)

type Recruitment struct {
	ID          int
	Title       string
	Place       string
	Start       time.Time
	End         time.Time
	Content     string
	Paid        bool
	Reward      string
	MemberLimit int
	UserID      int
	Type        RecruitmentType
	Tags        []int
}

package domain

import "time"

type RecruitmentType string

const (
	RecruitmentTypeDefault  RecruitmentType = "recruitment"
	RecruitmentTypeFreeTime RecruitmentType = "freeTime"
)

func (r RecruitmentType) String() string {
	switch r {
	case RecruitmentTypeDefault:
		return "recruitment"
	case RecruitmentTypeFreeTime:
		return "freeTime"
	}
	return "recruitment"
}

func ToRecruitmentType(t string) RecruitmentType {
	switch t {
	case "recruitment":
		return RecruitmentTypeDefault
	case "freeTime":
		return RecruitmentTypeFreeTime
	}
	return RecruitmentTypeDefault
}

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
	Created     time.Time
	Tags        []TagData
}

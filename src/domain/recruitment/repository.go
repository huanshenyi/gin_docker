package recruitment

import "gin_docker/src/domain"

type Repository interface {
	ListRecruitmentForUserID(tx domain.Tx, userID int) (Recruitments, error)
}

type Recruitments struct {
	Recruitments []domain.Recruitment
}

package recruitment

import "gin_docker/src/domain"

type Repository interface {
	ListRecruitmentForUserID(tx domain.Tx, userID int) (Recruitments, error)
	CreateRecruitment(tx domain.Tx, input domain.Recruitment) error
}

type Recruitments struct {
	Recruitments []domain.Recruitment
}

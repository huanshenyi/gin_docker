package recruitment

import "gin_docker/src/domain"

// 募集テーブルに対する操作
type Repository interface {
	ListRecruitmentForUserID(tx domain.Tx, userID int) (Recruitments, error)
	CreateRecruitment(tx domain.Tx, input domain.Recruitment) error //新規募集
}

type Recruitments struct {
	Recruitments []domain.Recruitment
}

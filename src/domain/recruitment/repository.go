package recruitment

import (
	"gin_docker/src/domain"
	"gin_docker/src/domain/user"
)

// Repository 募集テーブルに対する操作
type Repository interface {
	ListRecruitmentForUserID(tx domain.Tx, userID int) (Recruitments, error)                        // 自分がオーナーの募集を表示
	CreateRecruitment(tx domain.Tx, input domain.Recruitment) error                                 // 新規募集追加
	JoinListRecruitment(tx domain.Tx, userID int, page int, limit int) (JoinListRecruitment, error) // 自分が応募した募集を返す
}

type Recruitments struct {
	Recruitments []domain.Recruitment
}

type JoinListRecruitment struct {
	Recruitment []JoinRecruitment
	TotalPage   int
	TotalCount  int
}

type JoinRecruitment struct {
	Recruitment domain.Recruitment
	Owner       user.UserData
}

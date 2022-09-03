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
	GetRecruitmentByID(tx domain.Tx, id int) (domain.Recruitment, error)                            // ID指定された募集返す
	JoinRecruitment(tx domain.Tx, userID int, recruitmentID int) error                              // 募集参加
	CheckMemberLimit(tx domain.Tx, recruitmentID int, limit int) (bool, error)                      // CheckMemberLimit 募集満員かどうかチェック
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

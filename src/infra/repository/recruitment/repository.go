package recruitment

import (
	"gin_docker/src/domain"
	"gin_docker/src/domain/recruitment"
	"gin_docker/src/infra/model"
)

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) ListRecruitmentForUserID(tx domain.Tx, userID int) (recruitment.Recruitments, error) {
	conn := tx.DB()
	var rows []model.Recruitment
	if err := conn.Where("user_id = ?", userID).Find(&rows).Error; err != nil {
		return recruitment.Recruitments{}, err
	}
	rs := make([]domain.Recruitment, len(rows))
	for k, v := range rows {
		rs[k] = v.ToDomain()
	}
	return recruitment.Recruitments{
		Recruitments: rs,
	}, nil
}

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

func (r Repository) CreateRecruitment(tx domain.Tx, input domain.Recruitment) error {
	db := tx.DB()
	if err := db.Create(&model.Recruitment{
		Title:       input.Title,
		Place:       input.Place,
		Start:       input.Start,
		End:         input.End,
		Content:     input.Content,
		Paid:        input.Paid,
		Reward:      input.Reward,
		MemberLimit: input.MemberLimit,
		Type:        model.RecruitmentTypeToSQL(input.Type),
		UserID:      input.UserID,
	}).Error; err != nil {
		return err
	}
	return nil
}

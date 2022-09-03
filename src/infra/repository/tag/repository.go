package tag

import (
	"fmt"

	"gin_docker/src/domain"
	"gin_docker/src/infra/model"
)

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) List(tx domain.Tx, limit int, status int, keyWord string) ([]domain.TagData, error) {
	conn := tx.DB()
	var tags []model.Tag
	query := conn.Table(fmt.Sprintf("%s", new(model.Tag).TableName()))
	if keyWord != "" {
		query.Where("name LIKE ?", "%"+keyWord+"%")
	}
	if status != 0 {
		query.Where("status = ?", status)
	}
	if limit != 0 {
		query.Limit(limit)
	}
	if err := query.Find(&tags).Error; err != nil {
		return nil, err
	}
	tgList := make([]domain.TagData, len(tags))
	for i, v := range tags {
		tgList[i] = v.ToDomain()
	}
	return tgList, nil
}

func (r Repository) ListexistTags(tx domain.Tx, tagIDs []int) ([]domain.TagData, error) {
	conn := tx.ReadDB()
	var rows []model.Tag
	if err := conn.Select("id").Where("id IN (?)", tagIDs).Find(&rows).Error; err != nil {
		return nil, err
	}
	ids := make([]domain.TagData, len(rows))
	for k, i := range rows {
		ids[k].ID = i.ID
	}
	return ids, nil
}

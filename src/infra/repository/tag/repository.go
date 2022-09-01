package tag

import (
	"fmt"

	"gin_docker/src/domain"
	"gin_docker/src/domain/tag"
	"gin_docker/src/infra/model"
)

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) List(tx domain.Tx, limit int, status int, keyWord string) ([]tag.TagData, error) {
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
	tgList := make([]tag.TagData, len(tags))
	for i, v := range tags {
		tgList[i] = v.ToDomain()
	}
	return tgList, nil
}

func (r Repository) ListexistTags(tx domain.Tx, tagIDs []int) ([]int, error) {
	conn := tx.ReadDB()
	var rows []model.Tag
	if err := conn.Select("id").Where("id IN (?)", tagIDs).Find(&rows).Error; err != nil {
		return nil, err
	}
	ids := make([]int, len(rows))
	for k, i := range rows {
		ids[k] = i.ID
	}
	return ids, nil
}

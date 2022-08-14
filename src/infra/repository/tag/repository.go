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

func (r Repository) List(tx domain.Tx, limit int, status int) ([]tag.TagData, error) {
	conn := tx.DB()
	var tags []model.Tag
	query := conn.Table(fmt.Sprintf("%s", new(model.Tag).TableName()))
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

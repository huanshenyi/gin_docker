package tag

import "gin_docker/src/domain"

type Repository interface {
	List(tx domain.Tx, limit int, status int, keyWord string) ([]domain.TagData, error)
	ListexistTags(tx domain.Tx, tagIDs []int) ([]domain.TagData, error) // タグの存在チェック
}

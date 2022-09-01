package tag

import "gin_docker/src/domain"

type Repository interface {
	List(tx domain.Tx, limit int, status int, keyWord string) ([]TagData, error)
	ListexistTags(tx domain.Tx, tagIDs []int) ([]int, error) // タグの存在チェック
}

type TagData struct {
	ID     int
	Name   string
	Status int
}

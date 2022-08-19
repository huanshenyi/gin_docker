package tag

import "gin_docker/src/domain"

type Repository interface {
	List(tx domain.Tx, limit int, status int) ([]TagData, error)
}

type TagData struct {
	ID     int
	Name   string
	Status int
}
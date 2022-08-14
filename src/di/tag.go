package di

import (
	"gin_docker/src/controller"
	"gin_docker/src/domain"
	"gin_docker/src/infra/repository/tag"
	usecase "gin_docker/src/usecase/tag"
)

func newTagController(tx domain.Tx) controller.Tag {
	i := usecase.NewInteractor(
		tag.NewRepository(),
		tx,
	)
	return controller.Tag{Interactor: i}
}

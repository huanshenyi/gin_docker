package di

import (
	"gin_docker/src/controller"
	"gin_docker/src/domain"
	"gin_docker/src/infra/repository/user"
	usecase "gin_docker/src/usecase/user"
)

func newUserController(tx domain.Tx) controller.User {
	i := usecase.NewInteractor(
		user.NewRepository(),
		tx,
	)
	return controller.User{Interactor: i}
}

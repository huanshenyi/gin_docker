package di

import (
	"gin_docker/src/controller"
	"gin_docker/src/domain"
	"gin_docker/src/infra/repository/recruitment"
	usecase "gin_docker/src/usecase/recruitment"
)

func newRecruitmentController(tx domain.Tx) controller.Recruitment {
	i := usecase.NewInteractor(
		recruitment.NewRepository(),
		tx,
	)
	return controller.Recruitment{Interactor: i}
}

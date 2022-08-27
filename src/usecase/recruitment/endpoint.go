package recruitment

import (
	"gin_docker/src/domain"
	"gin_docker/src/domain/recruitment"
)

type Interactor interface {
	List(input ListInput) (ListOutput, error)
	Create(input CreateInput) error
	JoinList(input JoinListInput) (JoinListOutput, error)
	Join(input JoinInpt) error
}

func NewInteractor(repository recruitment.Repository, tx domain.Tx) Interactor {
	return &interactor{
		repository: repository,
		tx:         tx,
	}
}

type interactor struct {
	repository recruitment.Repository
	tx         domain.Tx
}

func (i *interactor) List(input ListInput) (output ListOutput, err error) {
	res, err := i.repository.ListRecruitmentForUserID(i.tx, input.UserID)
	if err != nil {
		return ListOutput{}, err
	}
	return ConvertRecruitmentOutput(res), nil
}

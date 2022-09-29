package recruitment

import (
	"gin_docker/src/domain"
	"gin_docker/src/domain/recruitment"
	"gin_docker/src/domain/tag"
)

type Interactor interface {
	List(input ListInput) (ListOutput, error)
	Create(input CreateInput) error
	JoinList(input JoinListInput) (JoinListOutput, error)
	Join(input JoinInpt) error
	PublicList(input PublicListInput) (PublicListOutput, error)
}

func NewInteractor(repository recruitment.Repository, tagRepository tag.Repository, tx domain.Tx) Interactor {
	return &interactor{
		repository:    repository,
		tagRepository: tagRepository,
		tx:            tx,
	}
}

type interactor struct {
	repository    recruitment.Repository
	tagRepository tag.Repository
	tx            domain.Tx
}

func (i *interactor) List(input ListInput) (output ListOutput, err error) {
	res, err := i.repository.ListRecruitmentForUserID(i.tx, input.UserID)
	if err != nil {
		return ListOutput{}, err
	}
	return ConvertRecruitmentOutput(res), nil
}

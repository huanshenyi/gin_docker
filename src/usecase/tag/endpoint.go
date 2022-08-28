package tag

import (
	"gin_docker/src/domain"
	"gin_docker/src/domain/tag"
)

type Interactor interface {
	List(input ListInput) ([]ListOutput, error)
}

func NewInteractor(repository tag.Repository, tx domain.Tx) Interactor {
	return &interactor{
		repository: repository,
		tx:         tx,
	}
}

type interactor struct {
	repository tag.Repository
	tx         domain.Tx
}

func (i *interactor) List(input ListInput) ([]ListOutput, error) {
	res, err := i.repository.List(i.tx, input.Limit, input.Status, input.Keyword)
	if err != nil {
		return nil, err
	}
	output := make([]ListOutput, len(res))
	for i, v := range res {
		output[i] = ListOutput{
			ID:     v.ID,
			Name:   v.Name,
			Status: v.Status,
		}
	}
	return output, nil
}

package user

import (
	"gin_docker/src/domain"
	"gin_docker/src/domain/user"
)

type Interactor interface {
	Regist(input RegistInput) error
	Login(input LoginInput) (UserToken, error)
}

func NewInteractor(repository user.Repository, tx domain.Tx) Interactor {
	return &interactor{
		repository: repository,
		tx:         tx,
	}
}

type interactor struct {
	repository user.Repository
	tx         domain.Tx
}

func (i *interactor) Regist(input RegistInput) error {
	err := i.repository.Regist(i.tx, user.RegistInput{
		Identfier: input.Identfier,
		Password:  input.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *interactor) Login(input LoginInput) (UserToken, error) {
	user, err := i.repository.Login(i.tx, user.LoginInput{
		Identfier:    input.Identfier,
		IdentityType: input.IdentityType,
		PassWord:     input.PassWord,
	})
	if err != nil {
		return UserToken{}, err
	}
	token, err := UserToToken(user)
	if err != nil {
		return UserToken{}, err
	}
	return UserToken{Token: token}, nil
}

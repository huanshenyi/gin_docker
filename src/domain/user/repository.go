package user

import "gin_docker/src/domain"

type Repository interface {
	Regist(tx domain.Tx, input RegistInput) error
	Login(tx domain.Tx, input LoginInput) (UserData, error)
}

type RegistInput struct {
	Identfier string
	Password  string
}

type LoginInput struct {
	Identfier    string //username | email | githubID
	IdentityType string //login_type
	PassWord     string
}

type UserData struct {
	UserName string
	Icon     string
}

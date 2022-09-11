package user

import "gin_docker/src/domain"

type Repository interface {
	Regist(tx domain.Tx, input RegistInput) error
	Login(tx domain.Tx, input LoginInput) (UserData, error)
	AddUserToken(tx domain.Tx, token string, userID int) error
	GetMyInfo(tx domain.Tx, userID int) (domain.UserProfile, error)
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
	ID       int
	UserName string
	Icon     string
}

func (u UserData) IsLoginedUser() bool {
	return u.ID != 0
}

func (u UserData) IsAnonymousUser() bool {
	return u.ID == 0
}

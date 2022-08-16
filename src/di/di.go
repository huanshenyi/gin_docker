package di

import (
	"gin_docker/src/controller"
	"gin_docker/src/domain"
	dblayer "gin_docker/src/infra/db"
)

// GssktService define GssktService struct
type GssktService struct {
	User          controller.User
	Tag           controller.Tag
	Recruitment   controller.Recruitment
	UserService   *dblayer.UserService
	ClientService ClientService
}

type ClientService struct {
}

// NewGssktService generate GssktService instance
func NewGssktService(tx domain.Tx) *GssktService {
	db := tx.DB()
	return &GssktService{
		User:        newUserController(tx),
		Tag:         newTagController(tx),
		Recruitment: newRecruitmentController(tx),
		UserService: dblayer.NewUserService(
			db,
			[]int{1, 2, 3},
		),
	}
}

package di

import (
	"gin_docker/src/controller"
	"gin_docker/src/domain"
)

// GssktService define GssktService struct
type GssktService struct {
	User          controller.User
	ClientService ClientService
}

type ClientService struct {
}

// NewGssktService generate GssktService instance
func NewGssktService(tx domain.Tx) *GssktService {
	return &GssktService{
		User: newUser(tx),
	}
}

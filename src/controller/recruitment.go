package controller

import (
	"gin_docker/src/domain/authenticator"
	"gin_docker/src/usecase/recruitment"
	"gin_docker/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Recruitment struct {
	Interactor recruitment.Interactor
}

func (t *Recruitment) List(c *gin.Context) {
	input, err := t.validateList(c)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	res, err := t.Interactor.List(input)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (t *Recruitment) validateList(c *gin.Context) (recruitment.ListInput, error) {
	user, ok := authenticator.GetUser(c)
	if !ok {
		return recruitment.ListInput{}, &utils.UnauthorizedError{Action: "recruitment List"}
	}
	ok = user.IsLoginedUser()
	if !ok {
		return recruitment.ListInput{}, &utils.UnauthorizedError{Action: "recruitment List"}
	}
	return recruitment.ListInput{UserID: user.ID}, nil
}

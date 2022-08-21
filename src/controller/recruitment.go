package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_docker/src/domain/authenticator"
	"gin_docker/src/usecase/recruitment"
	"gin_docker/src/utils"
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

func (t *Recruitment) Create(c *gin.Context) {
	input, err := t.validateCreateInput(c)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	err = t.Interactor.Create(input)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (t *Recruitment) validateCreateInput(c *gin.Context) (input recruitment.CreateInput, err error) {
	user, ok := authenticator.GetUser(c)
	if !ok {
		return recruitment.CreateInput{}, &utils.UnauthorizedError{Action: "recruitment create"}
	}
	ok = user.IsLoginedUser()
	if !ok {
		return recruitment.CreateInput{}, &utils.UnauthorizedError{Action: "recruitment create"}
	}
	input.UserID = user.ID
	err = c.BindJSON(&input)
	if err != nil {
		return recruitment.CreateInput{}, &utils.InvalidParamError{Err: err}
	}
	err = Validate(input)
	return
}

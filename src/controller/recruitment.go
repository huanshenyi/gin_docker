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

func (t *Recruitment) JoinList(c *gin.Context) {
	input, err := t.validateJoinListInput(c)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	res, err := t.Interactor.JoinList(input)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (t *Recruitment) validateJoinListInput(c *gin.Context) (input recruitment.JoinListInput, err error) {
	user, ok := authenticator.GetUser(c)
	if !ok {
		return recruitment.JoinListInput{}, &utils.UnauthorizedError{Action: "recruitment JoinList"}
	}
	ok = user.IsLoginedUser()
	if !ok {
		return recruitment.JoinListInput{}, &utils.UnauthorizedError{Action: "recruitment JoinList"}
	}
	input.UserID = user.ID
	err = c.Bind(&input)
	if err != nil {
		err = &utils.InvalidParamError{Err: err}
		return
	}
	err = Validate(input)
	return
}

// Join 募集に応募
func (t *Recruitment) Join(c *gin.Context) {
	input, err := t.validateJoinInput(c)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	err = t.Interactor.Join(input)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (t *Recruitment) validateJoinInput(c *gin.Context) (input recruitment.JoinInpt, err error) {
	user, ok := authenticator.GetUser(c)
	if !ok {
		return recruitment.JoinInpt{}, &utils.UnauthorizedError{Action: "recruitment Join"}
	}
	ok = user.IsLoginedUser()
	if !ok {
		return recruitment.JoinInpt{}, &utils.UnauthorizedError{Action: "recruitment Join"}
	}
	input.UserID = user.ID
	err = c.BindJSON(&input)
	if err != nil {
		err = &utils.InvalidParamError{Err: err}
		return
	}
	err = Validate(input)
	return
}

// PublicList - 公開募集リスト取得
func (t *Recruitment) PublicList(c *gin.Context) {
	input, err := t.validatePublicList(c)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	res, err := t.Interactor.PublicList(input)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (t *Recruitment) validatePublicList(c *gin.Context) (input recruitment.PublicListInput, err error) {
	err = c.BindQuery(&input)
	if err != nil {
		err = &utils.InvalidParamError{Err: err}
		return
	}
	err = Validate(input)

	return
}

func (t *Recruitment) Delete(c *gin.Context) {
	input, err := t.validateDeleteInput(c)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	err = t.Interactor.Delete(input)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (t *Recruitment) validateDeleteInput(c *gin.Context) (input recruitment.DeleteInput, err error) {
	user, ok := authenticator.GetUser(c)
	if !ok {
		return recruitment.DeleteInput{}, &utils.UnauthorizedError{Action: "recruitment delete"}
	}
	ok = user.IsLoginedUser()
	if !ok {
		return recruitment.DeleteInput{}, &utils.UnauthorizedError{Action: "recruitment delete"}
	}
	input.UserID = user.ID

	err = c.BindUri(&input)
	if err != nil {
		err = &utils.InvalidParamError{Err: err}
		return
	}
	return
}

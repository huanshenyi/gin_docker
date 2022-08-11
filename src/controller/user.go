package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_docker/src/usecase/user"
	"gin_docker/src/utils"
)

type User struct {
	Interactor user.Interactor
}

func (u *User) Regist(c *gin.Context) {
	input, err := u.validateRegistInput(c)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	err = u.Interactor.Regist(input)
	if err != nil {
		ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (u *User) validateRegistInput(c *gin.Context) (user.RegistInput, error) {
	var input user.RegistInput
	err := c.BindJSON(&input)
	if err != nil {
		return user.RegistInput{}, err
	}
	err = Validate(input)
	if err != nil {
		return user.RegistInput{}, err
	}
	return input, nil
}

func (u *User) Login(c *gin.Context) {
	input, err := u.validateLoginInput(c)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	token, err := u.Interactor.Login(input)
	if err != nil {
		ErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func (u *User) validateLoginInput(c *gin.Context) (user.LoginInput, error) {
	var input user.LoginInput
	err := c.BindJSON(&input)
	if err != nil {
		return user.LoginInput{}, &utils.InvalidParamError{Err: err}
	}
	err = Validate(input)
	if err != nil {
		return user.LoginInput{}, &utils.InvalidParamError{Err: err}
	}
	return input, nil
}

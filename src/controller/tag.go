package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_docker/src/usecase/tag"
)

type Tag struct {
	Interactor tag.Interactor
}

func (t *Tag) List(c *gin.Context) {
	input, err := t.validateListInput(c)
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

func (t *Tag) validateListInput(c *gin.Context) (tag.ListInput, error) {
	var input tag.ListInput
	err := c.BindQuery(&input)
	if err != nil {
		return tag.ListInput{}, err
	}
	return input, nil
}

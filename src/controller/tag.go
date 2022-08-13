package controller

import (
	"fmt"
	"gin_docker/src/domain/authenticator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func (t *Tag) List(c *gin.Context) {
	// t.validateListInput(c)
	user, _ := authenticator.GetUser(c)
	fmt.Println(user)
	fmt.Println(user.IsAnonymousUser())

	c.JSON(http.StatusOK, nil)
}

// func (t *Tag) validateListInput(c *gin.Context) {

// }

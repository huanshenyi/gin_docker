package authenticator

import (
	"github.com/gin-gonic/gin"

	"gin_docker/src/domain/user"
	"gin_docker/src/infra/db"
)

const userContextKey = "user"

type AuthedUser interface {
	IsLoginedUser() bool
	IsAnonymousUser() bool
}

func SetUser(ctx *gin.Context, user user.UserData) {
	ctx.Set(userContextKey, user)
}

func GetUser(ctx *gin.Context) (user.UserData, bool) {
	userData, ok := ctx.Get(userContextKey)
	if !ok {
		return user.UserData{}, false
	}
	return userData.(user.UserData), true
}

type Authenticator struct {
	userService *db.UserService
}

func NewAuthenticator(
	userService *db.UserService,
) *Authenticator {
	return &Authenticator{
		userService: userService,
	}
}

// gin.contextにユーザーデータ注入
func SetAuthedUser(ctx *gin.Context, user user.UserData) {
	// 今後ユーザータイプ増えればここの処理も増やす
	SetUser(ctx, user)
}

// ID 存在するかどうかで判断する
func (a *Authenticator) GetAnonymousUser() user.UserData {
	return user.UserData{ID: 0}
}

func (a *Authenticator) GetUserFromToken(token string, tokenType TokenType) (user.UserData, error) {
	if tokenType == UserToken {
		return a.userService.GetUserFromToken(token)
	}
	return user.UserData{ID: 0}, nil
}

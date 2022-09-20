package user

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"gin_docker/src/domain/user"
)

var mysecret = "mysecret"

type MyClaims struct {
	UserName string `json:"username"`
	Icon     string `json:"icon"`
	Group    string `json:"group"`
	jwt.RegisteredClaims
}

// UserDataをJWTTokenに変更
func UserToToken(user user.UserData) (string, error) {
	token, err := makeToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func makeToken(user user.UserData) (string, error) {
	claim := MyClaims{
		UserName: user.UserName,
		Icon:     user.Icon,
		Group:    user.Group.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()), //発行時間
			NotBefore: jwt.NewNumericDate(time.Now()), //発効時間
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(mysecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

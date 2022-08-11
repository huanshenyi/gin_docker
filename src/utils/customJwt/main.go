package customjwt

import (
	"errors"
	"gin_docker/src/usecase/user"

	"github.com/golang-jwt/jwt/v4"
)

var MySecret = "mysecret"

func Secret() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	}
}

func ParseUserToken(tokens string) (*user.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokens, &user.MyClaims{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that`s not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn`t handel this token")
			}
		}
	}
	if claims, ok := token.Claims.(*user.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn`t handel this token")
}

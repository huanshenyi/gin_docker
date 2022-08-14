package authenticator

import (
	"strings"

	"gin_docker/src/utils"
)

type TokenType int

const (
	UnKnownToken TokenType = iota
	UserToken
	OAuthToken
	ClientToken
)

// headerからtoken取得
// X-Token: user xxxxxxxxx
func GetTokenFromHeader(header string) (token string, tokenType TokenType, err error) {
	tokens := strings.Split(header, ",")
	for _, v := range tokens {
		v = strings.TrimSpace(v)
		kv := strings.Split(v, " ")
		if len(kv) != 2 {
			continue
		}
		if kv[0] == "User" {
			token = kv[1]
			tokenType = UserToken
			return
		} else if kv[0] == "OAuth" {
			token = kv[1]
			tokenType = OAuthToken
			return
		} else if kv[0] == "Client" {
			token = kv[1]
			tokenType = ClientToken
			return
		}
	}
	err = &utils.InvalidTokenError{}
	return
}

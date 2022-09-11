package controller

import (
	"testing"

	"gin_docker/src/osho"
	"gin_docker/src/usecase/user"

	domainUser "gin_docker/src/domain/user"
	"gin_docker/src/utils"
)

func Test_validateGetMyInfoInput(t *testing.T) {
	tests := []struct {
		name    string
		params  osho.HTTPParams
		want    user.GetMyInfoInput
		wantErr error
	}{
		{
			name: "ログインなしでは使えない",
			params: osho.HTTPParams{
				User: &domainUser.UserData{ID: 0},
			},
			wantErr: &utils.UnauthorizedError{Action: "get my info input"},
		},
		{
			name: "ログインすれば使える",
			params: osho.HTTPParams{
				User: &domainUser.UserData{ID: 1},
			},
			wantErr: nil,
			want:    user.GetMyInfoInput{UserID: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := osho.GetGinContext(tt.params)
			if err != nil {
				t.Fatal(err)
			}
			b := User{}
			got, err := b.validateGetMyInfoInput(c)
			osho.TestChecker(t, tt.want, tt.wantErr, got, err)
		})
	}
}

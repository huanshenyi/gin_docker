package controller

import (
	"gin_docker/src/domain/user"
	"gin_docker/src/osho"
	"gin_docker/src/usecase/recruitment"
	"gin_docker/src/utils"
	"testing"
)

func Test_validateList(t *testing.T) {
	tests := []struct {
		name    string
		params  osho.HTTPParams
		want    recruitment.ListInput
		wantErr error
	}{
		{
			name: "ログインしないと使えない",
			params: osho.HTTPParams{
				User: &user.UserData{ID: 0},
			},
			wantErr: &utils.UnauthorizedError{Action: "recruitment List"},
		},
		{
			name: "ログインすれば使える",
			params: osho.HTTPParams{
				User: &user.UserData{ID: 1},
			},
			wantErr: nil,
			want:    recruitment.ListInput{UserID: 1},
		},
	}
	for _, tt := range tests {
		c, err := osho.GetGinContext(tt.params)
		if err != nil {
			t.Fatal(err)
		}
		b := Recruitment{}
		got, err := b.validateList(c)
		osho.TestChecker(t, tt.want, tt.wantErr, got, err)
	}
}

package controller

import (
	"fmt"
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

func Test_validateUpdateMyInfoInput(t *testing.T) {
	tests := []struct {
		name    string
		params  osho.HTTPParams
		want    user.UpdateMyInfoInput
		wantErr error
	}{
		{
			name: "ログインなしでは使えない",
			params: osho.HTTPParams{
				User: &domainUser.UserData{ID: 0},
			},
			wantErr: &utils.UnauthorizedError{Action: "update my info input"},
		},
		{
			name: "ログインすれば使える",
			params: osho.HTTPParams{
				User: &domainUser.UserData{ID: 1},
				Body: map[string]interface{}{
					"userName": "newname",
					"icon":     "icon-url",
					"email":    "xx@xemil.com",
					"age":      18,
				},
			},
			wantErr: nil,
			want: user.UpdateMyInfoInput{
				UserID:     1,
				UserName:   "newname",
				Icon:       "icon-url",
				Email:      "xx@xemil.com",
				Sex:        0,
				LivingArea: "",
				Age:        18,
				Appeal:     "",
				Profession: "",
			},
		},
		{
			name: "sexは012以外の使用できない",
			params: osho.HTTPParams{
				User: &domainUser.UserData{ID: 1},
				Body: map[string]interface{}{
					"userName": "newname",
					"icon":     "icon-url",
					"sex":      3,
				},
			},
			wantErr: &utils.InvalidParamError{Err: fmt.Errorf("Key: 'UpdateMyInfoInput.Sex' Error:Field validation for 'Sex' failed on the 'oneof' tag")},
			want:    user.UpdateMyInfoInput{},
		},
		{
			name: "ageは1と120の間しか設定できない",
			params: osho.HTTPParams{
				User: &domainUser.UserData{ID: 1},
				Body: map[string]interface{}{
					"userName": "newname",
					"icon":     "icon-url",
					"age":      -1,
				},
			},
			wantErr: &utils.InvalidParamError{Err: fmt.Errorf("Key: 'UpdateMyInfoInput.Age' Error:Field validation for 'Age' failed on the 'gte' tag")},
			want:    user.UpdateMyInfoInput{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := osho.GetGinContext(tt.params)
			if err != nil {
				t.Fatal(err)
			}
			b := User{}
			got, err := b.validateUpdateMyInfoInput(c)
			osho.TestChecker(t, tt.want, tt.wantErr, got, err)
		})
	}
}

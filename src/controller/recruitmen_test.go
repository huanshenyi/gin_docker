package controller

import (
	"testing"
	"time"

	"gin_docker/src/domain"
	"gin_docker/src/domain/user"
	"gin_docker/src/osho"
	"gin_docker/src/usecase/recruitment"
	"gin_docker/src/utils"
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
		t.Run(tt.name, func(t *testing.T) {
			c, err := osho.GetGinContext(tt.params)
			if err != nil {
				t.Fatal(err)
			}
			b := Recruitment{}
			got, err := b.validateList(c)
			osho.TestChecker(t, tt.want, tt.wantErr, got, err)
		})
	}
}

func Test_validateCreateInput(t *testing.T) {
	startTime, _ := time.Parse(time.RFC3339Nano, "2022-09-01T00:00:00+09:00")
	endTime, _ := time.Parse(time.RFC3339Nano, "2022-09-01T00:00:00+09:00")
	tests := []struct {
		name    string
		params  osho.HTTPParams
		want    recruitment.CreateInput
		wantErr error
	}{
		{
			name: "ログインなしでは使えない",
			params: osho.HTTPParams{
				User: &user.UserData{ID: 0},
			},
			wantErr: &utils.UnauthorizedError{Action: "recruitment create"},
		},
		{
			name: "正常",
			params: osho.HTTPParams{
				User: &user.UserData{ID: 1},
				Body: map[string]interface{}{
					"title":       "タイトル",
					"place":       "場所",
					"start":       "2022-09-01T00:00:00+09:00",
					"end":         "2022-09-01T00:00:00+09:00",
					"content":     "内容",
					"paid":        true,
					"reward":      "100円",
					"memberLimit": 2,
					"type":        "recruitment",
				},
			},
			want: recruitment.CreateInput{
				UserID:      1,
				Title:       "タイトル",
				Place:       "場所",
				Start:       startTime,
				End:         endTime,
				Content:     "内容",
				Paid:        true,
				Reward:      "100円",
				MemberLimit: 2,
				Type:        domain.RecruitmentTypeDefault,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := osho.GetGinContext(tt.params)
			if err != nil {
				t.Fatal(err)
			}
			b := Recruitment{}
			got, err := b.validateCreateInput(c)
			osho.TestChecker(t, tt.want, tt.wantErr, got, err)
		})
	}
}

func Test_validateJoinListInput(t *testing.T) {
	tests := []struct {
		name    string
		params  osho.HTTPParams
		want    recruitment.JoinListInput
		wantErr error
	}{
		{
			name: "ログインなしでは使えない",
			params: osho.HTTPParams{
				User: &user.UserData{ID: 0},
			},
			wantErr: &utils.UnauthorizedError{Action: "recruitment JoinList"},
		},
		{
			name: "ログインすれば使える",
			params: osho.HTTPParams{
				User: &user.UserData{ID: 1},
			},
			wantErr: nil,
			want:    recruitment.JoinListInput{UserID: 1, Limit: 10, Page: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := osho.GetGinContext(tt.params)
			if err != nil {
				t.Fatal(err)
			}
			b := Recruitment{}
			got, err := b.validateJoinListInput(c)
			osho.TestChecker(t, tt.want, tt.wantErr, got, err)
		})
	}
}

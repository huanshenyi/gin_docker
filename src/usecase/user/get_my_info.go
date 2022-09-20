package user

import (
	"time"

	"gin_docker/src/domain"
)

type GetMyInfoInput struct {
	UserID int
}

type GetMyInfoOutput struct {
	ID         int       `json:"id"`
	UserName   string    `json:"username"`
	Icon       string    `json:"icon"`
	Email      string    `json:"email"`
	Sex        string    `json:"sex"`
	LivingArea string    `json:"livingArea"`
	Age        int       `json:"age"`
	Appeal     string    `json:"appeal"`
	Profession string    `json:"profession"`
	Group      string    `json:"group"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (i *interactor) GetMyInfo(input GetMyInfoInput) (GetMyInfoOutput, error) {
	res, err := i.repository.GetMyInfo(i.tx, input.UserID)
	if err != nil {
		return GetMyInfoOutput{}, err
	}
	return convertProfileOutput(res), nil
}

func convertProfileOutput(input domain.UserProfile) GetMyInfoOutput {
	return GetMyInfoOutput{
		ID:         input.UserID,
		UserName:   input.UserName,
		Icon:       input.Icon,
		Email:      input.Email,
		Sex:        input.Sex.String(),
		LivingArea: input.LivingArea,
		Age:        input.Age,
		Appeal:     input.Appeal,
		Profession: input.Profession,
		Group:      input.Group.String(),
		UpdatedAt:  input.UpdatedAt,
	}
}

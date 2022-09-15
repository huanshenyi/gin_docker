package user

import "gin_docker/src/domain"

type UpdateMyInfoInput struct {
	UserID     int
	UserName   string `json:"userName" validate:"required"`
	Icon       string `json:"icon" validate:"required"`
	Email      string `json:"email"`
	Sex        int    `json:"sex" validate:"oneof= 0 1 2"`
	LivingArea string `json:"livingArea"`
	Age        int    `json:"age" validate:"gte=0,lt=120"`
	Appeal     string `json:"appeal"`
	Profession string `json:"profession"`
}

func (u *UpdateMyInfoInput) toDomain() domain.UserProfile {
	return domain.UserProfile{
		UserID:     u.UserID,
		UserName:   u.UserName,
		Icon:       u.Icon,
		Email:      u.Email,
		Sex:        domain.SexType(u.Sex),
		LivingArea: u.LivingArea,
		Age:        u.Age,
		Appeal:     u.Appeal,
		Profession: u.Profession,
	}
}

func (i *interactor) UpdateMyInfo(input UpdateMyInfoInput) error {
	err := i.repository.UpdateMyInfo(i.tx, input.toDomain())
	if err != nil {
		return err
	}
	return nil
}

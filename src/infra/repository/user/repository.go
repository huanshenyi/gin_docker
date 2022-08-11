package user

import (
	"fmt"
	"gin_docker/src/domain"
	"gin_docker/src/domain/user"
	"gin_docker/src/infra/model"
	"math/rand"
)

type Repository struct {
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) Regist(tx domain.Tx, input user.RegistInput) error {
	defaultName := fmt.Sprintf("defaultName%d", rand.Intn(9999))
	row := model.User{UserName: defaultName, Icon: "https://avatars.githubusercontent.com/u/3580607?s=40&v=4",
		UserAuths: []model.UserAuth{{IdentityType: "default", Identfier: input.Identfier, Credential: input.Password}}}
	if err := tx.DB().Create(&row).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) Login(tx domain.Tx, input user.LoginInput) (user.UserData, error) {
	conn := tx.DB()
	var userData model.User
	query := conn.Table(fmt.Sprintf("%s as U", new(model.User).TableName())).
		Joins(fmt.Sprintf("LEFT JOIN %s as UA ON UA.user_id = U.id", new(model.UserAuth).TableName())).
		Where("UA.identity_type = ?", input.IdentityType).
		Where("UA.Identfier = ?", input.Identfier).
		Where("UA.Credential = ?", input.PassWord)
	if err := query.Find(&userData).Error; err != nil {
		return user.UserData{}, err
	}
	return user.UserData{
		UserName: userData.UserName,
		Icon:     userData.Icon,
	}, nil
}

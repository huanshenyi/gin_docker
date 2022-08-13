package user

import (
	"fmt"
	"gin_docker/src/domain"
	"gin_docker/src/domain/user"
	"gin_docker/src/infra/model"
	"math/rand"

	"gorm.io/gorm"
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
		ID:       userData.ID,
		UserName: userData.UserName,
		Icon:     userData.Icon,
	}, nil
}

func (r Repository) AddUserToken(tx domain.Tx, token string, userID int) error {
	var accessToken model.AccessToken
	result := tx.DB().Table(new(model.AccessToken).TableName()).Where("user_id = ?", userID).First(&accessToken)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			result := tx.DB().Table(new(model.AccessToken).TableName()).Create(&model.AccessToken{
				UserID:      userID,
				AccessToken: token,
			})
			if result.Error != nil {
				return result.Error
			}
		}
	} else {
		accessToken.AccessToken = token
		result := tx.DB().Table(new(model.AccessToken).TableName()).Where("user_id = ?", userID).Update("access_token", token)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

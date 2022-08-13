package db

import (
	"gorm.io/gorm"

	"gin_docker/src/domain/user"
	"gin_docker/src/infra/model"
	"gin_docker/src/utils"
)

func NewUserService(db *gorm.DB, demoUserIDs []int) *UserService {
	return &UserService{db: db, demoUserIDs: demoUserIDs}
}

type UserService struct {
	db          *gorm.DB
	demoUserIDs []int
}

// tokenからユーザーを探す
// 見つからない場合はUserNotFoundを返す
func (u *UserService) GetUserFromToken(token string) (user.UserData, error) {
	user, err := u.findUserFromToken(token)
	switch {
	case err == nil:
		return user, nil
	case err != nil:
		return u.fail(&utils.DBInternalError{Err: err})
	}
	return u.fail(&utils.DBInternalError{Err: err})
}

func (u *UserService) findUserFromToken(token string) (user.UserData, error) {
	db := u.db
	var tokenRow model.AccessToken
	if err := db.Where(&model.AccessToken{AccessToken: token}).First(&tokenRow).Error; err != nil {
		return u.fail(err)
	}
	userRow := model.User{ID: tokenRow.UserID}
	if err := db.First(&userRow).Error; err != nil {
		return u.fail(err)
	}
	return user.UserData{ID: userRow.ID}, nil
}

func (u *UserService) fail(err error) (user.UserData, error) {
	return user.UserData{}, err
}

package user

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"gin_docker/src/domain"
	"gin_docker/src/domain/user"
	"gin_docker/src/infra/model"
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
		Group:    domain.UserGroup(userData.Group),
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

func (r Repository) GetMyInfo(tx domain.Tx, userID int) (domain.UserProfile, error) {
	var row userProfile
	db := tx.DB()
	query := db.Select(`
	U.*,
	UP.email AS email,
	UP.sex AS sex,
	UP.living_area AS living_area,
	UP.age AS age,
	UP.appeal AS appeal,
	UP.profession AS profession
	`).
		Table(fmt.Sprintf("%s as U", new(model.User).TableName())).
		Joins(fmt.Sprintf("LEFT JOIN %s AS UP ON UP.user_id = U.id", new(model.UserProfile).TableName())).
		Where("U.id = ?", userID)

	if err := query.First(&row).Error; err != nil {
		return domain.UserProfile{}, err
	}

	return row.toDomain(), nil
}

type userProfile struct {
	ID         int       `gorm:"column:id"`
	UserName   string    `gorm:"column:username"`
	Icon       string    `gorm:"column:icon"`
	Email      string    `gorm:"column:email"`
	Sex        int       `gorm:"column:sex"`
	LivingArea string    `gorm:"column:living_area"`
	Age        int       `gorm:"column:age"`
	Appeal     string    `gorm:"column:appeal"`
	Profession string    `gorm:"column:profession"`
	Group      int       `gorm:"column:group"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (u userProfile) toDomain() domain.UserProfile {
	return domain.UserProfile{
		UserID:     u.ID,
		UserName:   u.UserName,
		Icon:       u.Icon,
		Email:      u.Email,
		Sex:        domain.SexType(u.Sex),
		LivingArea: u.LivingArea,
		Age:        u.Age,
		Appeal:     u.Appeal,
		Profession: u.Profession,
		Group:      domain.UserGroup(u.Group),
		UpdatedAt:  u.UpdatedAt,
	}
}

// TODO: 今のクエリいつか修正したい
func (r Repository) UpdateMyInfo(tx domain.Tx, user domain.UserProfile) error {
	db := tx.DB()

	if err := db.Model(model.User{}).Where("id", user.UserID).Updates(model.User{
		UserName: user.UserName,
		Icon:     user.Icon,
	}).Error; err != nil {
		return err
	}

	if err := db.Where("user_id", user.UserID).Updates(model.UserProfile{
		Email:      user.Email,
		Sex:        int(user.Sex),
		LivingArea: user.LivingArea,
		Age:        user.Age,
		Appeal:     user.Appeal,
		Profession: user.Profession,
	}).Error; err != nil {
		return err
	}

	return nil
}

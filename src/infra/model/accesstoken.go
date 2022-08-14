package model

import "time"

type AccessToken struct {
	ID          int       `gorm:"column:id"`           // id
	UserID      int       `gorm:"column:user_id"`      // user_id
	AccessToken string    `gorm:"column:access_token"` // access_token
	CreatedAt   time.Time `gorm:"column:created"`      // created
}

func (t *AccessToken) TableName() string {
	return "accesstokens"
}

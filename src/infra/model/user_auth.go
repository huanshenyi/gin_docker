package model

type UserAuth struct {
	ID           int    `gorm:"column:id"`
	UserID       int    `gorm:"column:user_id"`
	IdentityType string `gorm:"column:identity_type"` //phone | github タイプ
	Identfier    string `gorm:"column:identfier"`     // 公開の番号
	Credential   string `gorm:"credential"`           //token
}

func (m *UserAuth) TableName() string {
	return "user_auths"
}

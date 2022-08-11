package user

type LoginInput struct {
	Identfier    string `json:"identfier" form:"identfier" validate:"required"`       //username | email | githubID
	IdentityType string `json:"identityType" form:"identityType" validate:"required"` //login_type
	PassWord     string `json:"passWord" form:"passWord"`
}

type RegistInput struct {
	Identfier string `json:"identfier" form:"identfier" validate:"required"`
	Password  string `json:"passWord" form:"passWord" validate:"required"`
}

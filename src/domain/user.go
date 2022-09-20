package domain

import "time"

type PlatformType int

const (
	UndefinedPlatform PlatformType = iota
	IOS
	Android
	Web
)

type Device struct {
	Platform PlatformType
}

func (pt PlatformType) String() string {
	switch pt {
	case IOS:
		return "ios"
	case Android:
		return "android"
	case Web:
		return "web"
	}
	return "undefined"
}

type SexType int

const (
	Unknown SexType = iota
	Male
	Female
)

func (s SexType) String() string {
	switch s {
	case Unknown:
		return "unknown"
	case Male:
		return "male"
	case Female:
		return "female"
	}
	return "undefined"
}

type UserGroup int

const (
	GroupUnknown UserGroup = iota
	GroupAdmin
	GroupUser
)

func (s UserGroup) String() string {
	switch s {
	case GroupUnknown:
		return "unknown"
	case GroupAdmin:
		return "admin"
	case GroupUser:
		return "user"
	}
	return "unknown"
}

type UserProfile struct {
	UserID     int
	UserName   string
	Icon       string
	Email      string
	Sex        SexType
	LivingArea string
	Age        int
	Appeal     string
	Profession string
	Group      UserGroup
	UpdatedAt  time.Time
}

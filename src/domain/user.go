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
	UpdatedAt  time.Time
}

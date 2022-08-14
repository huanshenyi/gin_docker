package domain

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

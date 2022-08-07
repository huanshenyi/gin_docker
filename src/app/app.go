package app

type EnvType string

const (
	EnvDevelopment EnvType = "development"
	EnvProto       EnvType = "proto"
	EnvTesting     EnvType = "testing"
	EnvStaging     EnvType = "staging"
	EnvProduction  EnvType = "production"
	EnvTest        EnvType = "test"
)

func (e EnvType) String() string {
	return string(e)
}

var env = EnvDevelopment

func Env() EnvType {
	return env
}

func SetEnv(e EnvType) {
	env = e
}

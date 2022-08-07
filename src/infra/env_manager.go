package infra

import "os"

var EnvMan envManager

type envManager struct {
	AppEnv string `validate:"required"`

	DBUser     string `validate:"required"`
	DBPassword string `validate:"required"`
	DBHost     string `validate:"required"`
	DBPort     string `validate:"required"`
	DBName     string `validate:"required"`
	HOST       string `validate:"required"`
}

func init() {
	EnvMan = newEnvManager()
}

func newEnvManager() envManager {
	var em envManager
	em = envManager{
		AppEnv:     os.Getenv("APP_ENV"),
		DBUser:     os.Getenv("UB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		HOST:       os.Getenv("HOST"),
	}
	return em
}

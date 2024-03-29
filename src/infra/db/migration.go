package db

import (
	"gin_docker/src/app"
	"gin_docker/src/domain"
	"gin_docker/src/infra/model"
	"gin_docker/src/log_source"
)

func MigrateOnlyLocal(tx domain.Tx) {
	db := tx.DB()
	doMigrate := app.Env() == app.EnvDevelopment
	if !doMigrate {
		return
	}
	log_source.Log.Info("start migration...")

	hasTable := db.Migrator().HasTable("user_profiles")
	if hasTable {
		return
	}

	tables := []interface{}{
		// &model.User{},
		// &model.UserAuth{},
		// &model.Tag{},
		// &model.AccessToken{},
		// &model.Recruitment{},
		// &model.UserRecruitment{},
		// &model.RecruitmentTag{},
		&model.UserProfile{},
	}
	if err := db.Migrator().CreateTable(tables...); err != nil {
		log_source.Log.Error(err.Error())
	}
}

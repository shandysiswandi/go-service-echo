package welcomes

import (
	"go-rest-echo/db"
	"go-rest-echo/service"
)

type usecase struct {
	database *db.Database
	service  *service.Service
}

// NewUsecase is
func NewUsecase(db *db.Database, service *service.Service) Usecase {
	return &usecase{db, service}
}

func (u *usecase) CheckServiceJWT() bool {
	check := u.service.JWT
	if check == nil {
		return false
	}

	return true
}

func (u *usecase) CheckServiceLogger() bool {
	check := u.service.Logger
	if check == nil {
		return false
	}

	return true
}

func (u *usecase) CheckServiceSentry() bool {
	check := u.service.Sentry
	if check == nil {
		return false
	}

	return true
}

func (u *usecase) CheckServiceRedis() bool {
	check := u.service.Redis
	if check == nil {
		return false
	}

	return true
}

func (u *usecase) CheckDatabaseMysql() bool {
	check := u.database.Mysql
	if check == nil {
		return false
	}

	return true
}

func (u *usecase) CheckDatabasePostgresql() bool {
	check := u.database.Postgresql
	if check == nil {
		return false
	}

	return true
}

func (u *usecase) CheckDatabaseMongo() bool {
	check := u.database.Mongo
	if check == nil {
		return false
	}

	return true
}

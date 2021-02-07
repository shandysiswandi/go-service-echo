package welcomes

import (
	"go-rest-echo/db"
	"go-rest-echo/external"
	"go-rest-echo/external/jsonplaceholder.typicode.com"
	"go-rest-echo/service"
)

// Usecase is
type Usecase struct {
	database *db.Database
	service  *service.Service
	external *external.External
}

// NewUsecase is
func NewUsecase(db *db.Database, service *service.Service, external *external.External) *Usecase {
	return &Usecase{db, service, external}
}

// CheckServiceJWT is
func (u *Usecase) CheckServiceJWT() bool {
	check := u.service.JWT
	if check == nil {
		return false
	}

	return true
}

// CheckServiceSentry is
func (u *Usecase) CheckServiceSentry() bool {
	check := u.service.Sentry
	if check == nil {
		return false
	}

	return true
}

// CheckServiceRedis is
func (u *Usecase) CheckServiceRedis() bool {
	check := u.service.Redis
	if check == nil {
		return false
	}

	return true
}

// CheckDatabaseMysql is
func (u *Usecase) CheckDatabaseMysql() bool {
	check := u.database.Mysql
	if check == nil {
		return false
	}

	return true
}

// CheckDatabasePostgresql is
func (u *Usecase) CheckDatabasePostgresql() bool {
	check := u.database.Postgresql
	if check == nil {
		return false
	}

	return true
}

// CheckDatabaseMongo is
func (u *Usecase) CheckDatabaseMongo() bool {
	check := u.database.Mongo
	if check == nil {
		return false
	}

	return true
}

// CheckExternalJSONPlaceHolder is map[string]interface{}
func (u *Usecase) CheckExternalJSONPlaceHolder() (map[string]interface{}, error) {
	fetch, err := u.external.JSONPlaceHolder.FetchPost()
	if err != nil {
		return nil, err
	}

	get, err := u.external.JSONPlaceHolder.GetPost(1)
	if err != nil {
		return nil, err
	}

	pCreate := jsonplaceholder.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	create, err := u.external.JSONPlaceHolder.CreatePost(pCreate)
	if err != nil {
		return nil, err
	}

	pUpdate := jsonplaceholder.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	update, err := u.external.JSONPlaceHolder.UpdatePost(pUpdate, 1)
	if err != nil {
		return nil, err
	}

	if err = u.external.JSONPlaceHolder.DeletePost(1); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"fetch_post":  fetch[1:2],
		"get_post":    get,
		"create_post": create,
		"update_post": update,
		"delete_post": true,
	}, nil
}

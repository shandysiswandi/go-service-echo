package welcomes

import (
	"go-rest-echo/app/library/jwtlib"
	"go-rest-echo/app/library/redislib"
	"go-rest-echo/app/library/sentrylib"
	"go-rest-echo/db"
	"go-rest-echo/external/jsonplaceholder"
)

// Usecase is
type Usecase struct {
	database        *db.Database
	jwt             *jwtlib.JWT
	redis           *redislib.Redis
	sentry          *sentrylib.Sentry
	jsonPlaceHolder *jsonplaceholder.JSONPlaceHolder
}

// NewUsecase is
func NewUsecase(db *db.Database, j *jwtlib.JWT, r *redislib.Redis, s *sentrylib.Sentry, jph *jsonplaceholder.JSONPlaceHolder) *Usecase {
	return &Usecase{db, j, r, s, jph}
}

// CheckLibraryJWT is
func (u *Usecase) CheckLibraryJWT() bool {
	return u.jwt != nil
}

// CheckLibrarySentry is
func (u *Usecase) CheckLibrarySentry() bool {
	return u.sentry != nil
}

// CheckLibraryRedis is
func (u *Usecase) CheckLibraryRedis() bool {
	return u.redis != nil
}

// CheckDatabaseMysql is
func (u *Usecase) CheckDatabaseMysql() bool {
	return u.database.Mysql != nil
}

// CheckDatabasePostgresql is
func (u *Usecase) CheckDatabasePostgresql() bool {
	return u.database.Postgresql != nil
}

// CheckDatabaseMongo is
func (u *Usecase) CheckDatabaseMongo() bool {
	return u.database.Mongo != nil
}

// CheckExternalJSONPlaceHolder is map[string]interface{}
func (u *Usecase) CheckExternalJSONPlaceHolder() (map[string]interface{}, error) {
	fetch, err := u.jsonPlaceHolder.FetchPost()
	if err != nil {
		return nil, err
	}

	get, err := u.jsonPlaceHolder.GetPost(1)
	if err != nil {
		return nil, err
	}

	pCreate := jsonplaceholder.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	create, err := u.jsonPlaceHolder.CreatePost(pCreate)
	if err != nil {
		return nil, err
	}

	pUpdate := jsonplaceholder.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	update, err := u.jsonPlaceHolder.UpdatePost(pUpdate, 1)
	if err != nil {
		return nil, err
	}

	if err = u.jsonPlaceHolder.DeletePost(1); err != nil {
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

package route

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomeRoute is
func HomeRoute(e *echo.Echo, config *config.Config, db *db.Database) {
	e.GET("/", func(c echo.Context) error {
		mysqlDB := "connection ok"
		if db.Mysql == nil {
			mysqlDB = "no connection"
		} else {
			d, err := db.Mysql.DB()
			if err != nil {
				mysqlDB = "connection ok, but there is some problem"
			}
			if err = d.Ping(); err != nil {
				mysqlDB = "connection ok, but can't PING the connection"
			}
		}

		postgresqlDB := "connection ok"
		if db.Posrgresql == nil {
			postgresqlDB = "no connection"
		} else {
			d, err := db.Posrgresql.DB()
			if err != nil {
				mysqlDB = "connection ok, but there is some problem"
			}
			if err = d.Ping(); err != nil {
				postgresqlDB = "connection ok, but can't PING the connection"
			}
		}

		mongoDB := "connection ok"

		var response struct {
			Env                          string   `json:"env"`
			Port                         string   `json:"port"`
			Name                         string   `json:"name"`
			DatabaseDriver               []string `json:"database_driver"`
			DatabaseConnectionMysql      string   `json:"database_connection_mysql"`
			DatabaseConnectionPostgresql string   `json:"database_connection_postgresql"`
			DatabaseConnectionMongo      string   `json:"database_connection_mongo"`
		}

		response.Env = config.App.Env
		response.Port = config.App.Port
		response.Name = config.App.Name
		response.DatabaseDriver = config.SchemaDatabases
		response.DatabaseConnectionMysql = mysqlDB
		response.DatabaseConnectionPostgresql = postgresqlDB
		response.DatabaseConnectionMongo = mongoDB

		return c.JSON(http.StatusOK, response)
	})
}

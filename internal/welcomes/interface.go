package welcomes

import "github.com/labstack/echo/v4"

type (
	// WelcomeDelivery is
	WelcomeDelivery interface {
		Home(echo.Context) error
		MonitorDatabase(echo.Context) error
		MonitorService(echo.Context) error
	}

	// Usecase is
	Usecase interface {
		CheckServiceJWT() bool
		CheckServiceLogger() bool
		CheckServiceSentry() bool
		CheckServiceRedis() bool

		CheckDatabaseMysql() bool
		CheckDatabasePostgresql() bool
		CheckDatabaseMongo() bool
	}
)

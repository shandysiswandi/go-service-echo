package welcomes

type usecase struct{}

// NewUsecase is
func NewUsecase() Usecase {
	return &usecase{}
}

func (usecase) CheckServiceSentry() bool {

	return true
}

func (usecase) CheckServiceRedis() bool {

	return true
}

func (usecase) CheckDatabaseMysql() bool {

	return true
}

func (usecase) CheckDatabasePostgresql() bool {

	return true
}

func (usecase) CheckDatabaseMongo() bool {

	return true
}

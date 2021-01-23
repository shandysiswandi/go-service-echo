package user

import (
	"go-rest-echo/db"
	"go-rest-echo/entity"
	"go-rest-echo/usecase/user"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

// NewMysql is contstructor
func NewMysql(db *db.Database) user.Repository {
	return &mysqlRepository{db: db.Mysql}
}

func (m *mysqlRepository) Fetch() (t []*entity.User, err error) {
	err = m.db.Find(t).Error
	if err != nil {
		return nil, err
	}

	return t, err
}

func (m *mysqlRepository) Get(string) (*entity.User, error) {
	return nil, nil
}

func (m *mysqlRepository) Create(*entity.User) error {
	return nil
}

func (m *mysqlRepository) Update(*entity.User, string) error {
	return nil
}

func (m *mysqlRepository) Delete(string) error {
	return nil
}

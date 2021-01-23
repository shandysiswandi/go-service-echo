package task

import (
	"go-rest-echo/db"
	"go-rest-echo/entity"
	"go-rest-echo/usecase/task"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

// NewMysql is contstructor
func NewMysql(db *db.Database) task.Repository {
	return &mysqlRepository{db: db.Mysql}
}

func (m *mysqlRepository) Fetch() (t []*entity.Task, err error) {
	return nil, nil
}

func (m *mysqlRepository) Get(string) (*entity.Task, error) {
	return nil, nil
}

func (m *mysqlRepository) Create(*entity.Task) error {
	return nil
}

func (m *mysqlRepository) Update(*entity.Task, string) error {
	return nil
}

func (m *mysqlRepository) Delete(string) error {
	return nil
}

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
func NewMysql() task.Repository {
	return &mysqlRepository{db: db.GetMysqlDB()}
}

func (m *mysqlRepository) Fetch() (t []*entity.Task, err error) {
	err = m.db.Find(t).Error
	if err != nil {
		return nil, err
	}

	return t, err
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

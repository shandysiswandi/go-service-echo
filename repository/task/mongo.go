package task

import (
	"go-rest-echo/db"
	"go-rest-echo/entity"
	"go-rest-echo/usecase/task"

	"gorm.io/gorm"
)

type mongoRepository struct {
	db *gorm.DB
}

// NewMongo is contstructor
func NewMongo() task.Repository {
	return &mongoRepository{db: db.GetMongoDB()}
}

func (m *mongoRepository) Fetch() (t []*entity.Task, err error) {
	err = m.db.Find(t).Error
	if err != nil {
		return nil, err
	}

	return t, err
}

func (m *mongoRepository) Get(string) (*entity.Task, error) {
	return nil, nil
}

func (m *mongoRepository) Create(*entity.Task) error {
	return nil
}

func (m *mongoRepository) Update(*entity.Task, string) error {
	return nil
}

func (m *mongoRepository) Delete(string) error {
	return nil
}

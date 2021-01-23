package task

import (
	"go-rest-echo/db"
	"go-rest-echo/entity"
	"go-rest-echo/usecase/task"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

// NewMongo is contstructor
func NewMongo(db *db.Database) task.Repository {
	return &mongoRepository{db: db.Mongo}
}

func (m *mongoRepository) Fetch() (t []*entity.Task, err error) {
	return nil, nil
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

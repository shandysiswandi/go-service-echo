package repository

import (
	"go-rest-echo/db"
	"go-rest-echo/internal/blogs"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

// NewMongo is contstructor
func NewMongo(db *db.Database) blogs.BlogRepository {
	return &mongoRepository{db: db.Mongo}
}

func (m *mongoRepository) Fetch() (*[]blogs.Blog, error) {
	return nil, nil
}

func (m *mongoRepository) Get(ID string) (*blogs.Blog, error) {
	return nil, nil
}

func (m *mongoRepository) Create(b *blogs.Blog) error {
	return nil
}

func (m *mongoRepository) Update(b *blogs.Blog, ID string) error {

	return nil
}

func (m *mongoRepository) Delete(ID string) error {

	return nil
}

package blogs

import (
	"go-rest-echo/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

// NewMongo is contstructor
func NewMongo(db *db.Database) Repository {
	return &mongoRepository{db: db.Mongo}
}

func (m *mongoRepository) Fetch() ([]Blog, error) {
	return nil, nil
}

func (m *mongoRepository) Get(ID string) (*Blog, error) {
	return nil, nil
}

func (m *mongoRepository) Create(b *Blog) error {
	return nil
}

func (m *mongoRepository) Update(b *Blog, ID string) error {

	return nil
}

func (m *mongoRepository) UpdateField(b *Blog, ID string) error {

	return nil
}

func (m *mongoRepository) Delete(ID string) error {

	return nil
}

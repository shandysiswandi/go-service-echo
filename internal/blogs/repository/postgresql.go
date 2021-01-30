package repository

import (
	"go-rest-echo/db"
	"go-rest-echo/internal/blogs"

	"gorm.io/gorm"
)

type postgresqlRepository struct {
	db *gorm.DB
}

// NewPostgresql is contstructor
func NewPostgresql(db *db.Database) blogs.BlogRepository {
	return &postgresqlRepository{db: db.Postgresql}
}

func (m *postgresqlRepository) Fetch() (*[]blogs.Blog, error) {
	t := new([]blogs.Blog)

	if err := m.db.Find(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *postgresqlRepository) Get(ID string) (*blogs.Blog, error) {
	t := new(blogs.Blog)

	if err := m.db.First(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *postgresqlRepository) Create(b *blogs.Blog) error {
	if err := m.db.Create(b).Error; err != nil {
		return err
	}

	return nil
}

func (m *postgresqlRepository) Update(b *blogs.Blog, ID string) error {
	model := blogs.Blog{ID: ID}
	q := m.db.Model(&model).Updates(b)

	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *postgresqlRepository) Delete(ID string) error {
	model := new(blogs.Blog)
	q := m.db.Delete(model, ID)

	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

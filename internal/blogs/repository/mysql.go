package repository

import (
	"go-rest-echo/db"
	"go-rest-echo/internal/blogs"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

// NewMysql is contstructor
func NewMysql(db *db.Database) blogs.BlogRepository {
	return &mysqlRepository{db: db.Mysql}
}

func (m *mysqlRepository) Fetch() (*[]blogs.Blog, error) {
	t := new([]blogs.Blog)

	if err := m.db.Find(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *mysqlRepository) Get(ID string) (*blogs.Blog, error) {
	t := new(blogs.Blog)

	if err := m.db.First(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *mysqlRepository) Create(b *blogs.Blog) error {
	if err := m.db.Create(b).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlRepository) Update(b *blogs.Blog, ID string) error {
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

func (m *mysqlRepository) Delete(ID string) error {
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

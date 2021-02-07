package blogs

import (
	"go-rest-echo/db"

	"gorm.io/gorm"
)

type postgresqlRepository struct {
	db *gorm.DB
}

// NewPostgresql is contstructor
func NewPostgresql(db *db.Database) BlogRepository {
	return &postgresqlRepository{db: db.Postgresql}
}

func (m *postgresqlRepository) Fetch() (*[]Blog, error) {
	t := new([]Blog)

	if err := m.db.Find(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *postgresqlRepository) Get(ID string) (*Blog, error) {
	t := new(Blog)

	if err := m.db.First(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *postgresqlRepository) Create(b *Blog) error {
	if err := m.db.Create(b).Error; err != nil {
		return err
	}

	return nil
}

func (m *postgresqlRepository) Update(b *Blog, ID string) error {
	model := Blog{ID: ID}
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
	model := new(Blog)
	q := m.db.Delete(model, ID)

	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

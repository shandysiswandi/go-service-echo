package blogs

import (
	"go-rest-echo/db"

	"gorm.io/gorm"
)

type postgresqlRepository struct {
	db *gorm.DB
}

// NewPostgresql is contstructor
func NewPostgresql(db *db.Database) Repository {
	return &postgresqlRepository{db.SQL}
}

func (m *postgresqlRepository) Fetch() ([]Blog, error) {
	b := []Blog{}

	if err := m.db.Find(&b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func (m *postgresqlRepository) Get(ID string) (*Blog, error) {
	b := &Blog{ID: ID}

	if err := m.db.First(&b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func (m *postgresqlRepository) Create(b *Blog) error {
	if err := m.db.Create(b).Error; err != nil {
		return err
	}

	return nil
}

func (m *postgresqlRepository) Update(b *Blog, ID string) error {
	q := m.db.Model(&Blog{ID: ID}).Updates(b)

	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *postgresqlRepository) UpdateField(b *Blog, ID string) error {
	q := m.db.Model(&Blog{ID: ID}).Updates(b)

	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *postgresqlRepository) Delete(ID string) error {
	q := m.db.Delete(&Blog{ID: ID})

	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

package blogs

import (
	"go-rest-echo/db"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

// NewMysql is contstructor
func NewMysql(db *db.Database) Repository {
	return &mysqlRepository{db: db.Mysql}
}

func (m *mysqlRepository) Fetch() (*[]Blog, error) {
	t := new([]Blog)

	if err := m.db.Find(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *mysqlRepository) Get(ID string) (*Blog, error) {
	t := new(Blog)

	if err := m.db.First(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *mysqlRepository) Create(b *Blog) error {
	if err := m.db.Create(b).Error; err != nil {
		return err
	}

	return nil
}

func (m *mysqlRepository) Update(b BlogPayloadPut, ID string) error {
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

func (m *mysqlRepository) UpdateField(b BlogPayloadPatch, ID string) error {
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

func (m *mysqlRepository) Delete(ID string) error {
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

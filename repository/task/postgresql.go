package task

import (
	"go-rest-echo/db"
	"go-rest-echo/entity"
	"go-rest-echo/usecase/task"

	"gorm.io/gorm"
)

type postgresqlRepository struct {
	db *gorm.DB
}

// NewPostgresql is contstructor
func NewPostgresql(db *db.Database) task.Repository {
	return &postgresqlRepository{db: db.Posrgresql}
}

func (m *postgresqlRepository) Fetch() (t []*entity.Task, err error) {
	err = m.db.Find(t).Error
	if err != nil {
		return nil, err
	}

	return t, err
}

func (m *postgresqlRepository) Get(string) (*entity.Task, error) {
	return nil, nil
}

func (m *postgresqlRepository) Create(*entity.Task) error {
	return nil
}

func (m *postgresqlRepository) Update(*entity.Task, string) error {
	return nil
}

func (m *postgresqlRepository) Delete(string) error {
	return nil
}

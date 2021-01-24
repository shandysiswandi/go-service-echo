package tasks

import (
	"go-rest-echo/db"
	"log"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

// NewMysql is contstructor
func NewMysql(db *db.Database) TaskRepository {
	return &mysqlRepository{db: db.Mysql}
}

func (m *mysqlRepository) Fetch() (*[]Task, error) {
	t := new([]Task)

	if err := m.db.Find(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *mysqlRepository) Get(ID string) (*Task, error) {
	t := new(Task)

	if err := m.db.First(t).Error; err != nil {
		return nil, err
	}

	return t, nil
}

func (m *mysqlRepository) Create(t *Task) (*string, error) {
	q := m.db.Create(t)
	if q.Error != nil {
		return nil, q.Error
	}
	log.Println(q)
	return nil, nil
}

func (m *mysqlRepository) Update(*Task, string) error {
	return nil
}

func (m *mysqlRepository) Delete(string) error {
	return nil
}

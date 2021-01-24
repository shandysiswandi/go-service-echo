package users

import (
	"go-rest-echo/db"
	"log"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

// NewMysql is contstructor
func NewMysql(db *db.Database) UserRepository {
	return &mysqlRepository{db: db.Mysql}
}

func (m *mysqlRepository) Fetch() (*[]User, error) {
	u := new([]User)

	if err := m.db.Find(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (m *mysqlRepository) Get(ID string) (*User, error) {
	u := new(User)

	if err := m.db.First(u, "id", ID).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (m *mysqlRepository) Create(u *User) error {
	if err := m.db.Create(u).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *mysqlRepository) Update(*User, string) error {
	return nil
}

func (m *mysqlRepository) Delete(string) error {
	return nil
}

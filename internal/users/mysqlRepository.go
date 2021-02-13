package users

import (
	"go-rest-echo/db"

	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

// NewMysql is contstructor
func NewMysql(db *db.Database) UserRepository {
	return &mysqlRepository{db: db.SQL}
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

func (m *mysqlRepository) GetByEmail(email string) (*User, error) {
	u := new(User)

	if err := m.db.First(u, "email", email).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (m *mysqlRepository) Create(u *User) error {
	q := m.db.Create(u)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (m *mysqlRepository) Update(u *User, ID string) error {
	q := m.db.Where("id = ?", ID).Updates(u)
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *mysqlRepository) Delete(ID string) error {
	u := new(User)

	q := m.db.Delete(u, "id", ID)
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

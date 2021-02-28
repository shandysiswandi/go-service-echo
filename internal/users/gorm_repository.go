package users

import (
	"go-service-echo/db"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

// NewGormRepository is contstructor
func NewGormRepository(db *db.Database) UserRepository {
	return &gormRepository{db: db.SQL}
}

func (m *gormRepository) Fetch() (*[]User, error) {
	u := new([]User)

	if err := m.db.Find(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (m *gormRepository) Get(ID string) (*User, error) {
	u := new(User)

	if err := m.db.First(u, "id", ID).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (m *gormRepository) GetByEmail(email string) (*User, error) {
	u := new(User)

	if err := m.db.First(u, "email", email).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (m *gormRepository) Create(u *User) error {
	q := m.db.Create(u)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (m *gormRepository) Update(u *User, ID string) error {
	q := m.db.Where("id = ?", ID).Updates(u)
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (m *gormRepository) Delete(ID string) error {
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

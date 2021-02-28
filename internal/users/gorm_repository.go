package users

import (
	"go-service-echo/db"
	"time"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

// NewGormRepository is contstructor
func NewGormRepository(db *db.Database) UserRepository {
	return &gormRepository{db: db.SQL}
}

func (m *gormRepository) Fetch() (Users, error) {
	var us Users
	return us, m.db.Find(&us).Error
}

func (m *gormRepository) Get(ID string) (*User, error) {
	u := &User{}
	return u, m.db.First(u, "id", ID).Error
}

func (m *gormRepository) GetByEmail(email string) (*User, error) {
	u := &User{}
	return u, m.db.First(u, "email", email).Error
}

func (m *gormRepository) Create(u *User) error {
	return m.db.Create(u).Error
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
	q := m.db.Debug().Model(&User{}).Where("id = ? AND deleted_at IS NULL", ID).Update("deleted_at", time.Now())
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

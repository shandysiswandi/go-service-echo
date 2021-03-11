package users

import (
	"time"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

// NewGormRepository is contstructor
func NewGormRepository(db *gorm.DB) UserRepository {
	return &gormRepository{db: db}
}

func (m *gormRepository) Fetch() (Users, error) {
	var us Users
	var table = User{}.TableName()

	return us, m.db.Table(table).Find(&us).Error
}

func (m *gormRepository) Get(ID string) (*User, error) {
	var u = new(User)
	var table = User{}.TableName()

	if err := m.db.Table(table).First(u, "id", ID).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (m *gormRepository) GetByEmail(email string) (*User, error) {
	var u = new(User)
	var table = User{}.TableName()

	if err := m.db.Table(table).First(u, "email", email).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (m *gormRepository) Create(u *UserCreatePayload) error {
	var table = User{}.TableName()

	return m.db.Table(table).Create(u).Error
}

func (m *gormRepository) Update(u *UserUpdatePayload, ID string) error {
	var table = User{}.TableName()

	query := m.db.Table(table).Where("id = ?", ID).Updates(u)
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected < 1 {
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

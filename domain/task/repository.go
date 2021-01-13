package task

import (
	"go-rest-echo/db/mysql"
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// Repository is
type Repository interface {
	Fetch(u *[]entity.Task) error
	Get(u *entity.Task, id string) error
	Create(u *entity.Task) error
	Update(u *entity.Task, id string) error
	Delete(u *entity.Task, id string) error
}

// NewRepository is contstructor
func NewRepository() Repository {
	return &repository{DB: mysql.GetDB()}
}

func (r *repository) Fetch(u *[]entity.Task) error {
	query := r.DB.Find(u)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *repository) Get(u *entity.Task, id string) error {
	query := r.DB.Where("id = ?", id).First(u)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *repository) Create(u *entity.Task) error {
	query := r.DB.Create(u)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *repository) Update(u *entity.Task, id string) error {
	query := r.DB.Where("id = ?", id).Updates(u)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (r *repository) Delete(u *entity.Task, id string) error {
	query := r.DB.Where("id = ?", id).Delete(u)
	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

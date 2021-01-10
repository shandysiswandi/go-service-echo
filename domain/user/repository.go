package user

import (
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// Repository is
type Repository interface {
	Fetch(u *[]entity.User) *gorm.DB
	Get(u *entity.User, id string) *gorm.DB
	Create(u *entity.User) *gorm.DB
	Update(u *entity.User, id string) *gorm.DB
	Delete(u *entity.User, id string) *gorm.DB
}

// NewRepository is contstructor
func NewRepository() Repository {
	return &repository{DB: nil}
}

func (r *repository) Fetch(u *[]entity.User) *gorm.DB {
	return r.DB.Find(u)
}

func (r *repository) Get(u *entity.User, id string) *gorm.DB {
	return r.DB.Where("id = ?", id).First(u)
}

func (r *repository) Create(u *entity.User) *gorm.DB {
	return r.DB.Create(u)
}

func (r *repository) Update(u *entity.User, id string) *gorm.DB {
	return r.DB.Where("id = ?", id).Updates(u)
}

func (r *repository) Delete(u *entity.User, id string) *gorm.DB {
	return r.DB.Where("id = ?", id).Delete(u)
}

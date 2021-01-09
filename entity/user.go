package entity

import (
	"errors"
	"go-rest-echo/entity/base"
	"go-rest-echo/helper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserTable is
const UserTable string = "users"

// User is
type User struct {
	base.UUID
	Name     string `gorm:"type:varchar(100)" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(100)" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(100)" json:"password" validate:"required"`
	Task     []Task `gorm:"foreignKey:UserID" json:"tasks"`
	base.Timestamp
}

// BeforeCreate is
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	u.Password, err = helper.HashPassword(u.Password)
	if err != nil {
		err = errors.New("can't hash `user` password")
	}
	return nil
}

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
	Name     string `json:"name" gorm:"type:varchar(100); not null" validate:"required,min=5"`
	Email    string `json:"email" gorm:"type:varchar(100); unique" validate:"required,email,min=5"`
	Password string `json:"-" gorm:"type:varchar(100)" validate:"required,min=6"`
	Task     []Task `json:"tasks" gorm:"foreignKey:UserID"`
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

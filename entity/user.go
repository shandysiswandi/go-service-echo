package entity

import (
	"errors"
	"go-rest-echo/app/base"
	"go-rest-echo/helper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is
type User struct {
	base.UUID
	Name     string `json:"name" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email,min=5"`
	Password string `json:"password" validate:"required,min=6"`
	Task     []Task `json:"tasks"`
	base.Timestamp
}

/*
 * Hooks GORM
 */

// TableName is
func (User) TableName() string {
	return "users"
}

// BeforeCreate is
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	u.Password, err = helper.HashPassword(u.Password)
	if err != nil {
		return errors.New("can't hash `user` password")
	}
	return nil
}

// BeforeUpdate is
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.Password, err = helper.HashPassword(u.Password)
	if err != nil {
		return errors.New("can't hash `user` password")
	}
	return nil
}

package create

import (
	"go-rest-echo/database/mysql"
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

// Repository is
func Repository(user *entity.User) *gorm.DB {
	return mysql.DB.Create(user)
}

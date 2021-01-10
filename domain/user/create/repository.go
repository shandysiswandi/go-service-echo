package create

import (
	"go-rest-echo/database/mysql"
	"go-rest-echo/entity"

	"gorm.io/gorm"
)

// Repository is
func Repository(u *entity.User) *gorm.DB {
	return mysql.DB.Create(u)
}
